package wal

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"

	"github.com/bimal009/db/pkg/constants"
	apperr "github.com/bimal009/db/pkg/err"
)

type Record struct {
	Type  byte
	Key   []byte
	Value []byte
}

type WalRecord struct {
	checksum uint32
	bodySize uint32
	Type     byte
	keySize  uint32
	Key      []byte
	Value    []byte
}

var crcTable = crc32.MakeTable(crc32.Castagnoli)

func Encode(r *Record) ([]byte, error) {
	bodySize := constants.MinBodySize + len(r.Key) + len(r.Value)

	if bodySize > constants.MaxBodySize {
		return nil, fmt.Errorf("%w: body size %d exceeds max %d", apperr.ErrInvalidSize, bodySize, constants.MaxBodySize)
	}

	if len(r.Key) == 0 {
		return nil, fmt.Errorf("%w: empty key", apperr.ErrInvalidSize)
	}

	tmpbuf := make([]byte, bodySize+constants.HeaderSize)

	binary.LittleEndian.PutUint32(tmpbuf[constants.ChecksumLen:constants.HeaderSize], uint32(bodySize))
	tmpbuf[constants.HeaderSize] = byte(r.Type)
	binary.LittleEndian.PutUint32(tmpbuf[constants.HeaderSize+constants.TypeLen:constants.HeaderSize+constants.MinBodySize], uint32(len(r.Key)))
	copy(tmpbuf[constants.HeaderSize+constants.MinBodySize:], r.Key)
	copy(tmpbuf[constants.HeaderSize+constants.MinBodySize+len(r.Key):], r.Value)

	hash := crc32.Checksum(tmpbuf[constants.ChecksumLen:], crcTable)
	binary.LittleEndian.PutUint32(tmpbuf[:constants.ChecksumLen], hash)
	return tmpbuf, nil
}
