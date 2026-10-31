package wal

import (
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"
	"testing"

	"github.com/bimal009/db/pkg/constants"
)

func TestEncode(t *testing.T) {
	record := &Record{
		Type:  constants.WalPut,
		Key:   []byte("name"),
		Value: []byte("john"),
	}

	encoded, err := Encode(record)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	expectedSize := constants.HeaderSize + constants.MinBodySize + len(record.Key) + len(record.Value)

	if len(encoded) != expectedSize {
		t.Fatalf("expected size %d, got %d", expectedSize, len(encoded))
	}

	storedCRC := binary.LittleEndian.Uint32(encoded[:constants.ChecksumLen])
	actualCRC := crc32.Checksum(encoded[constants.ChecksumLen:], crcTable)

	if storedCRC != actualCRC {
		t.Fatalf("checksum mismatch: expected %d, got %d", actualCRC, storedCRC)
	}

	t.Logf("WAL record:\n%s", hex.Dump(encoded))
	t.Log("WAL encoding successful")
}
