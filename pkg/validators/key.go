package validators

import (
	"fmt"

	"github.com/bimal009/db/pkg/err"
)

func ValidateKey(keyLength int) error {
	if keyLength == 0 {
		return fmt.Errorf("%w: empty key", err.ErrInvalidSize)
	}

	return nil
}
