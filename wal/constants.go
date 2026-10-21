package wal

// Record layout on disk (all integers are little-endian):
//
// [checksum 4][body size 4][type 1][key size 4][key][value]
//
// The header contains the checksum and body size.
// The body contains the type, key size, key, and value.
// The checksum is calculated over the body only.
// The body size does not include the header.

const (
	ChecksumSize = 4 // Bytes used to store the checksum
	BodySize     = 4 // Bytes used to store the body size
	TypeSize     = 1 // Bytes used to store the record type
	KeySize      = 4 // Bytes used to store the key size
)

const (
	HeaderSize  = ChecksumSize + BodySize
	MinBodySize = TypeSize + KeySize
)

const (
	MaxRecordSize = 16 * 1024 * 1024 // 16 MB maximum record size
	MaxBodySize   = MaxRecordSize - HeaderSize
)

const (
	WalFileExtension = ".wal"
	WalDirPath       = "../db/volume/wal"
	FileNumberDigits = 8
)
