package constants

// Entry layout on disk, shared by WAL records and SSTable entries
// (all integers are little-endian):
//
// [checksum 4][body size 4][type 1][key size 4][key][value]
//
// The header contains the checksum and body size.
// The body contains the type, key size, key, and value.
// The checksum is calculated over the body size and the body (everything after the checksum field).
// The body size does not include the header.
//
// SSTable file layout:
//
// [data section: entries][index section: [key size 4][key][offset 8] ...][footer 40]
//
// Footer layout:
//
// [index offset 8][index size 8][entry count 8][index checksum 4][magic 8][footer checksum 4]
// The footer checksum is calculated over the first 36 bytes of the footer.

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
	MaxRecordSize        = 1 << 20
	MaxBodySize          = MaxRecordSize - HeaderSize
	MaxSegmentSize       = 64 << 20
	MaxMemtableSizeBytes = 4 << 20 // flush when the memtable reaches 4 MB
)

const (
	TypeOff    = HeaderSize              // 8
	KeySizeOff = TypeOff + TypeLen       // 9
	KeyOff     = KeySizeOff + KeySizeLen // 13
)
const (
	OffsetLen         = 8
	IndexOffsetLen    = 8
	IndexSizeLen      = 8
	EntryCountLen     = 8
	IndexChecksumLen  = 4
	MagicLen          = 8
	FooterChecksumLen = 4

	FooterSize     = IndexOffsetLen + IndexSizeLen + EntryCountLen + IndexChecksumLen + MagicLen + FooterChecksumLen // 40
	FooterBodySize = FooterSize - FooterChecksumLen                                                                  // 36, the bytes the footer checksum covers

	SSTMagic uint64 = 0x5353543030303031 // "SST00001", the last digits act as the version
)

const (
	WalFileExtension = ".wal"
	WalDirPath       = "../db/volume/wal"
	SSTFileExtension = ".sst"
	SSTDirPath       = "../db/volume/sst"
	FileNumberDigits = 8
)

// Entry types, used as the type byte in both WAL records and SSTable entries.
const (
	Put        byte = iota + 1 // 1
	Delete                     // 2, an SSTable tombstone is a Delete entry with an empty value
	BatchBegin                 // 3, reserved
	BatchEnd                   // 4, reserved
)
