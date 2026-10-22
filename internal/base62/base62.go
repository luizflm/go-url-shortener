package base62

import (
	"errors"
	"math"
	"strings"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func ConvertToBase62(decimal int64) string {
	if decimal == 0 {
		return "0"
	}

	b62 := ""

	numerator := decimal
	for numerator > 0 {
		remainder := numerator % 62
		b62 = string(base62Chars[remainder]) + b62
		numerator = numerator / 62
	}

	return b62
}

func ConvertToDecimal(base62 string) (int64, error) {
	if base62 == "0" {
		return 0, nil
	}

	var result int64

	for _, c := range base62 {
		idx := strings.IndexRune(base62Chars, c)

		if idx < 0 {
			return -1, errors.New("invalid base62 encoding")
		}

		digit := int64(idx)
		if result > (math.MaxInt64-digit)/62 {
			return -1, errors.New("base62 value overflows int64")
		}

		result = result*62 + digit
	}
	return result, nil
}
