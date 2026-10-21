package wal

func ValidateMaxBodySize(body []byte) error {
	if len(body) > MaxBodySize {
		return ErrInvalidSize
	}

	return nil
}
