package engine

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"

	"github.com/bimal009/db/pkg/constants"
	"github.com/bimal009/db/pkg/err"
)

type SSTRecord struct {
	Deleted bool
	Key     []byte
	Value   []byte
}

type SSTIndexRecord struct {
	Key    string
	Offset uint64
}

type SSTFooter struct {
	IndexOffset   uint64
	IndexSize     uint64
	EntryCount    uint64
	IndexChecksum uint32
}

var crcTable = crc32.MakeTable(crc32.Castagnoli)

// [checksum 4][body size 4][type 1][key size 4][key][value] -> data
// [key size: 4][key][offset: 8] -> indexed
func EncodeSSTRecord(r *SSTRecord) ([]byte, error) {
	keyLen := len(r.Key)

	if keyLen == 0 {
		return nil, fmt.Errorf("%w: empty key", err.ErrInvalidSize)
	}

	bodySize := constants.MinBodySize + keyLen + len(r.Value)
	if bodySize > constants.MaxBodySize {
		return nil, fmt.Errorf("%w: body size %d exceeds max %d", err.ErrInvalidSize, bodySize, constants.MaxBodySize)
	}

	typ := constants.Put
	if r.Deleted {
		typ = constants.Delete
	}

	buf := make([]byte, constants.HeaderSize+bodySize)
	valueOff := constants.KeyOff + keyLen

	binary.LittleEndian.PutUint32(buf[constants.ChecksumLen:constants.HeaderSize], uint32(bodySize))
	buf[constants.TypeOff] = typ
	binary.LittleEndian.PutUint32(buf[constants.KeySizeOff:constants.KeyOff], uint32(keyLen))
	copy(buf[constants.KeyOff:], r.Key)
	copy(buf[valueOff:], r.Value)

	binary.LittleEndian.PutUint32(buf[:constants.ChecksumLen], crc32.Checksum(buf[constants.ChecksumLen:], crcTable))
	return buf, nil
}

func EncodeSSTFooter(r *SSTFooter) ([]byte, error)

func HandleIndex(r *SSTIndexRecord) ([]byte, error) {
	keyLen := len(r.Key)

	if keyLen == 0 {
		return nil, fmt.Errorf("%w: empty key", err.ErrInvalidSize)
	}

}
