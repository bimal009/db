package wal

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"

	"github.com/bimal009/db/pkg/constants"
	apperr "github.com/bimal009/db/pkg/err"
	"github.com/bimal009/db/pkg/validators"
)

type Record struct {
	Type  byte
	Key   []byte
	Value []byte
}

var crcTable = crc32.MakeTable(crc32.Castagnoli)

func EncodeWalRecord(r *Record) ([]byte, error) {
	keyLen := len(r.Key)

	if err := validators.ValidateKey(keyLen); err != nil {
		return nil, err
	}

	bodySize := constants.MinBodySize + keyLen + len(r.Value)
	if bodySize > constants.MaxBodySize {
		return nil, fmt.Errorf("%w: body size %d exceeds max %d", apperr.ErrInvalidSize, bodySize, constants.MaxBodySize)
	}

	buf := make([]byte, constants.HeaderSize+bodySize)
	valueOff := constants.KeyOff + keyLen

	binary.LittleEndian.PutUint32(buf[constants.ChecksumLen:constants.HeaderSize], uint32(bodySize))
	buf[constants.TypeOff] = r.Type
	binary.LittleEndian.PutUint32(buf[constants.KeySizeOff:constants.KeyOff], uint32(keyLen))
	copy(buf[constants.KeyOff:], r.Key)
	copy(buf[valueOff:], r.Value)

	binary.LittleEndian.PutUint32(buf[:constants.ChecksumLen], crc32.Checksum(buf[constants.ChecksumLen:], crcTable))
	return buf, nil
}
