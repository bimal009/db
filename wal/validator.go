package wal

import (
	"github.com/bimal009/db/pkg/constants"
	apperr "github.com/bimal009/db/pkg/err"
)

func ValidateMaxBodySize(body []byte) error {
	if len(body) > constants.MaxBodySize {
		return apperr.ErrInvalidSize
	}

	return nil
}
