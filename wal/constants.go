package wal

// Record layout on disk (all integers are little-endian):
//
// [checksum 4][body size 4][type 1][key size 4][key][value]

//note::
// checksum->[body size 4][type 1][key size 4][key][value]
// bodysize->[type 1][key size 4][key][value]

//
// The header contains the checksum and body size.
// The body contains the type, key size, key, and value.
// The checksum is calculated over the body only.
// The body size does not include the header.

const (
	ChecksumLen = 4
	BodySizeLen = 4
	TypeLen     = 1
	KeySizeLen  = 4
)

const (
	HeaderSize  = ChecksumLen + BodySizeLen
	MinBodySize = TypeLen + KeySizeLen
)

const (
	MaxRecordSize = 64 * 1024 * 1024 // 64 MB maximum record size
	MaxBodySize   = MaxRecordSize - HeaderSize
)

const (
	WalFileExtension = ".wal"
	WalDirPath       = "../db/volume/wal"
	FileNumberDigits = 8
)

const (
	WalPut        byte = iota + 1 // 1
	WalDelete                     // 2
	WalBatchBegin                 // 3, reserved
	WalBatchEnd                   // 4, reserved
)
