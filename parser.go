package gg

import (
	"strconv"
)

const (
	ErrorInvalidStringToIntConversion = ConstError("invalid string to int conversion")
)

func ParseInt(input string) (int, error) {
	integer, err := strconv.Atoi(input)
	if err != nil {
		return 0, ErrorInvalidStringToIntConversion
	}

	return integer, nil
}

func ParseUInt(input string) (uint, error) {
	integer, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return 0, ErrorInvalidStringToIntConversion
	}

	return uint(integer), nil
}
