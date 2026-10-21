package wal

import (
	"context"
	"hash/crc32"
)

var crcTable = crc32.MakeTable(crc32.Castagnoli)

type Hash interface {
	HashBody(ctx context.Context, body []byte) (uint32, error)
	CheckHash(ctx context.Context, checksum uint32, body []byte) error
}

func HashBody(ctx context.Context, body []byte) (uint32, error) {
	if err := ValidateMaxBodySize(body); err != nil {
		return 0, err
	}

	return crc32.Checksum(body, crcTable), nil
}

func CheckHash(ctx context.Context, checksum uint32, body []byte) error {
	calculated, err := HashBody(ctx, body)
	if err != nil {
		return err
	}

	if calculated != checksum {
		return ErrChecksumMismatch
	}

	return nil
}
