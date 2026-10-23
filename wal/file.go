package wal

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bimal009/db/pkg/utils"
)

func EnsureDir(dir string) error {
	if dir == "" {
		return ErrEmptyDirPath
	}

	info, err := os.Stat(dir)
	if err == nil {
		if !info.IsDir() {
			return fmt.Errorf("%w: %s", ErrNotADirectory, dir)
		}
		return nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("wal: stat %s: %w", dir, err)
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("%w: %s: %w", ErrFailedToCreateDir, dir, err)
	}
	return nil
}

// func FindLatestFile() (string, error) {

// }

func GetFiles(dir string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", ErrFilesNotFound, dir, err)
	}
	return entries, nil
}

func LatestFile(entries []os.DirEntry, ext string) (string, uint64, error) {
	var highest uint64
	var name string
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ext {
			continue
		}
		n, err := utils.ParseFileNumber(e.Name(), ext)
		if err != nil {
			return "", 0, err
		}
		if n > highest {
			highest = n
			name = e.Name()
		}
	}
	return name, highest, nil
}

func CreateWalFile(dir string) (*os.File, uint64, error) {
	entries, err := GetFiles(dir)
	if err != nil {
		return nil, 0, err
	}

	_, latest, err := LatestFile(entries, WalFileExtension)
	if err != nil {
		return nil, 0, err
	}

	next := latest + 1
	name := fmt.Sprintf("%0*d%s", FileNumberDigits, next, WalFileExtension)

	f, err := os.OpenFile(filepath.Join(dir, name),
		os.O_CREATE|os.O_EXCL|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, 0, err
	}
	return f, next, nil
}
func Append(dir string, record *Record) error {
	buf, err := Encode(record)
	if err != nil {
		return err
	}

	entries, err := GetFiles(dir)
	if err != nil {
		return err
	}

	name, _, err := LatestFile(entries, WalFileExtension)
	if err != nil {
		return err
	}

	var f *os.File
	if name == "" {
		f, _, err = CreateWalFile(dir)
		if err != nil {
			return err
		}
	} else {
		f, err = os.OpenFile(filepath.Join(dir, name), os.O_WRONLY|os.O_APPEND, 0o644)
		if err != nil {
			return fmt.Errorf("%w: %s: %w", ErrFailedToOpenFile, name, err)
		}
	}
	defer f.Close()

	if _, err := f.Write(buf); err != nil {
		return fmt.Errorf("%w: %s: %w", ErrFailedToWriteFile, f.Name(), err)
	}
	if err := f.Sync(); err != nil {
		return fmt.Errorf("%w: %s: %w", ErrFailedToSyncFile, f.Name(), err)
	}
	return nil
}
