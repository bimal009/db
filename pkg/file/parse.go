package file

import (
	"strconv"
	"strings"

	apperr "github.com/bimal009/db/pkg/err"
)

func ParseFileNumber(fileName, ext string) (uint64, error) {
	name, ok := strings.CutSuffix(fileName, ext)
	if !ok {
		return 0, apperr.ErrInvalidFileName
	}
	number, err := strconv.ParseUint(name, 10, 64)
	if err != nil {
		return 0, apperr.ErrInvalidFileName
	}
	return number, nil
}

func GetTheHighestNumber(highestRef *uint64, n uint64) {
	if *highestRef < n {
		*highestRef = n
	}
}
