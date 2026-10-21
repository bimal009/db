package utils

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidFileName = errors.New("utils: invalid file name")

func ParseFileNumber(fileName, ext string) (uint64, error) {
	name, ok := strings.CutSuffix(fileName, ext)
	if !ok {
		return 0, ErrInvalidFileName
	}
	number, err := strconv.ParseUint(name, 10, 64)
	if err != nil {
		return 0, ErrInvalidFileName
	}
	return number, nil
}
