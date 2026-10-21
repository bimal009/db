package wal

import "errors"

var (
	ErrIncompleteRecord   = errors.New("wal: incomplete record")
	ErrChecksumMismatch   = errors.New("wal: checksum mismatch")
	ErrRecordTooLarge     = errors.New("wal: record too large")
	ErrInvalidType        = errors.New("wal: invalid record type")
	ErrInvalidSize        = errors.New("wal: invalid record size")
	ErrWalDirNotFound     = errors.New("wal: directory not found")
	ErrEmptyDirPath       = errors.New("wal: directory path is empty")
	ErrNotADirectory      = errors.New("wal: path exists but is not a directory")
	ErrFailedToCreateDir  = errors.New("wal: failed to create directory")
	ErrFileAlreadyExists  = errors.New("wal: file already exists")
	ErrFailedToCreateFile = errors.New("wal: failed to create file")
)
