package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	apperr "github.com/bimal009/db/pkg/err"
)

func EnsureDir(dir string) error {
	if dir == "" {
		return apperr.ErrEmptyDirPath
	}

	info, err := os.Stat(dir)
	if err == nil {
		if !info.IsDir() {
			return fmt.Errorf("%w: %s", apperr.ErrNotADirectory, dir)
		}
		return nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("file: stat %s: %w", dir, err)
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("%w: %s: %w", apperr.ErrFailedToCreateDir, dir, err)
	}
	return nil
}

func GetFiles(dir string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", apperr.ErrFilesNotFound, dir, err)
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
		n, err := ParseFileNumber(e.Name(), ext)
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
