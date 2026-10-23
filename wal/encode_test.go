package wal

import (
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"
	"testing"
)

func TestEncode(t *testing.T) {
	record := &Record{
		Type:  WalPut,
		Key:   []byte("name"),
		Value: []byte("john"),
	}

	encoded, err := Encode(record)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	expectedSize := HeaderSize + MinBodySize + len(record.Key) + len(record.Value)

	if len(encoded) != expectedSize {
		t.Fatalf("expected size %d, got %d", expectedSize, len(encoded))
	}

	storedCRC := binary.LittleEndian.Uint32(encoded[:ChecksumLen])
	actualCRC := crc32.Checksum(encoded[ChecksumLen:], crcTable)

	if storedCRC != actualCRC {
		t.Fatalf("checksum mismatch: expected %d, got %d", actualCRC, storedCRC)
	}

	t.Logf("WAL record:\n%s", hex.Dump(encoded))
	t.Log("WAL encoding successful")
}
