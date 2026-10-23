package wal

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
)

type Record struct {
	Type  byte
	Key   []byte
	Value []byte
}

var crcTable = crc32.MakeTable(crc32.Castagnoli)

func Encode(r *Record) ([]byte, error) {
	bodySize := MinBodySize + len(r.Key) + len(r.Value)

	if bodySize > MaxBodySize {
		return nil, fmt.Errorf("%w: body size %d exceeds max %d", ErrInvalidSize, bodySize, MaxBodySize)
	}

	if len(r.Key) == 0 {
		return nil, fmt.Errorf("%w: empty key", ErrInvalidSize)
	}

	tmpbuf := make([]byte, bodySize+HeaderSize)

	binary.LittleEndian.PutUint32(tmpbuf[ChecksumLen:HeaderSize], uint32(bodySize))
	tmpbuf[HeaderSize] = byte(r.Type)
	binary.LittleEndian.PutUint32(tmpbuf[HeaderSize+TypeLen:HeaderSize+MinBodySize], uint32(len(r.Key)))
	copy(tmpbuf[HeaderSize+MinBodySize:], r.Key)
	copy(tmpbuf[HeaderSize+MinBodySize+len(r.Key):], r.Value)

	hash := crc32.Checksum(tmpbuf[ChecksumLen:], crcTable)
	binary.LittleEndian.PutUint32(tmpbuf[:ChecksumLen], hash)
	return tmpbuf, nil
}
