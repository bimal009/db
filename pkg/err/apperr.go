package err

import "errors"

var (
	ErrMemtableTooLarge = errors.New("memtable: too large")
	ErrKeyNotFound      = errors.New("memtable: key not found")
	ErrKeyDeleted       = errors.New("memtable: key deleted")

	ErrIncompleteRecord   = errors.New("wal: incomplete record")
	ErrChecksumMismatch   = errors.New("wal: checksum mismatch")
	ErrRecordTooLarge     = errors.New("wal: record too large")
	ErrFileTooLarge       = errors.New("wal: file too large")
	ErrInvalidType        = errors.New("wal: invalid record type")
	ErrInvalidSize        = errors.New("wal: invalid record size")
	ErrWalDirNotFound     = errors.New("wal: directory not found")
	ErrEmptyDirPath       = errors.New("wal: directory path is empty")
	ErrNotADirectory      = errors.New("wal: path exists but is not a directory")
	ErrFailedToCreateDir  = errors.New("wal: failed to create directory")
	ErrFileAlreadyExists  = errors.New("wal: file already exists")
	ErrFilesNotFound      = errors.New("wal: file doesnot exists")
	ErrFailedToCreateFile = errors.New("wal: failed to create file")
	ErrFailedToOpenFile   = errors.New("wal: failed to open file")
	ErrFailedToSyncFile   = errors.New("wal: failed to sync file")
	ErrFailedToWriteFile  = errors.New("wal: failed to write file")

	ErrInvalidFileName = errors.New("utils: invalid file name")
)
