package wal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bimal009/db/pkg/constants"
	apperr "github.com/bimal009/db/pkg/err"
	"github.com/bimal009/db/pkg/file"
)

func CreateWalFile(dir string) (*os.File, uint64, error) {
	entries, err := file.GetFiles(dir)
	if err != nil {
		return nil, 0, err
	}

	_, latest, err := file.LatestFile(entries, constants.WalFileExtension)
	if err != nil {
		return nil, 0, err
	}

	next := latest + 1
	name := fmt.Sprintf("%0*d%s", constants.FileNumberDigits, next, constants.WalFileExtension)

	f, err := os.OpenFile(filepath.Join(dir, name),
		os.O_CREATE|os.O_EXCL|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, 0, err
	}
	return f, next, nil
}

func Append(dir string, record *Record) error {
	buf, err := EncodeWalRecord(record)
	if err != nil {
		return err
	}

	entries, err := file.GetFiles(dir)
	if err != nil {
		return err
	}

	name, _, err := file.LatestFile(entries, constants.WalFileExtension)
	if err != nil {
		return err
	}

	var f *os.File
	if name != "" {
		f, err = os.OpenFile(filepath.Join(dir, name), os.O_WRONLY|os.O_APPEND, 0o644)
		if err != nil {
			return fmt.Errorf("%w: %s: %w", apperr.ErrFailedToOpenFile, name, err)
		}

		info, err := f.Stat()
		if err != nil {
			f.Close()
			return fmt.Errorf("wal: stat %s: %w", name, err)
		}
		if info.Size()+int64(len(buf)) > constants.MaxSegmentSize {
			f.Close()
			name = ""
		}
	}

	if name == "" {
		f, _, err = CreateWalFile(dir)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	if _, err := f.Write(buf); err != nil {
		return fmt.Errorf("%w: %s: %w", apperr.ErrFailedToWriteFile, f.Name(), err)
	}
	if err := f.Sync(); err != nil {
		return fmt.Errorf("%w: %s: %w", apperr.ErrFailedToSyncFile, f.Name(), err)
	}
	return nil
}
