package engine

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"

	"github.com/bimal009/db/pkg/constants"
	"github.com/bimal009/db/pkg/err"
	"github.com/bimal009/db/pkg/validators"
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
	if err := validators.ValidateKey(len(r.Key)); err != nil {
		return nil, err
	}

	bodySize := constants.MinBodySize + len(r.Key) + len(r.Value)
	if bodySize > constants.MaxBodySize {
		return nil, fmt.Errorf("%w: body size %d exceeds max %d", err.ErrInvalidSize, bodySize, constants.MaxBodySize)
	}

	typ := constants.Put
	if r.Deleted {
		typ = constants.Delete
	}

	buf := make([]byte, constants.HeaderSize+bodySize)
	valueOff := constants.KeyOff + len(r.Key)

	binary.LittleEndian.PutUint32(buf[constants.ChecksumLen:constants.HeaderSize], uint32(bodySize))
	buf[constants.TypeOff] = typ
	binary.LittleEndian.PutUint32(buf[constants.KeySizeOff:constants.KeyOff], uint32(len(r.Key)))
	copy(buf[constants.KeyOff:], r.Key)
	copy(buf[valueOff:], r.Value)

	binary.LittleEndian.PutUint32(buf[:constants.ChecksumLen], crc32.Checksum(buf[constants.ChecksumLen:], crcTable))
	return buf, nil
}

func EncodeSSTFooter(r *SSTFooter) []byte {
	buf := make([]byte, constants.FooterSize)

	binary.LittleEndian.PutUint64(buf[constants.FooterIndexOffsetOff:], r.IndexOffset)
	binary.LittleEndian.PutUint64(buf[constants.FooterIndexSizeOff:], r.IndexSize)
	binary.LittleEndian.PutUint64(buf[constants.FooterEntryCountOff:], r.EntryCount)
	binary.LittleEndian.PutUint32(buf[constants.FooterIndexChecksumOff:], r.IndexChecksum)
	binary.LittleEndian.PutUint64(buf[constants.FooterMagicOff:], constants.SSTMagic)

	binary.LittleEndian.PutUint32(buf[constants.FooterChecksumOff:], crc32.Checksum(buf[:constants.FooterChecksumOff], crcTable))
	return buf
}

func EncodeSSTIndexRecord(r *SSTIndexRecord) ([]byte, error) {
	keyLen := len(r.Key)

	if err := validators.ValidateKey(keyLen); err != nil {
		return nil, err
	}

	buf := make([]byte, constants.KeySizeLen+keyLen+constants.OffsetLen)

	binary.LittleEndian.PutUint32(buf[:constants.KeySizeLen], uint32(keyLen))
	copy(buf[constants.KeySizeLen:], r.Key)
	binary.LittleEndian.PutUint64(buf[constants.KeySizeLen+keyLen:], r.Offset)

	return buf, nil
}
