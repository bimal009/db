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
func CreateWalFile(dir string) (*os.File, uint64, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, 0, err
	}

	var highest uint64
	for _, e := range entries {
		n, err := utils.ParseFileNumber(e.Name(), WalFileExtension)
		if err == nil && n > highest {
			highest = n
		}
	}

	next := highest + 1
	name := fmt.Sprintf("%0*d%s", FileNumberDigits, next, WalFileExtension)

	f, err := os.OpenFile(filepath.Join(dir, name),
		os.O_CREATE|os.O_EXCL|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, 0, err
	}
	return f, next, nil
}
