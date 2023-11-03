package converter

import (
	"errors"
	"github.com/Krynegal/numeral-system-translator.git/internal/validators"
	"strings"
)

type Converter struct{}

func New() *Converter {
	return &Converter{}
}

func (cv Converter) Convert(num string, base int, toBase int) (string, error) {
	decNum, err := cv.ConvertFromAnyBaseToDecimal(num, base)
	if err != nil {
		return "", err
	}

	result, err := cv.ConvertFromDecimalToBaseX(decNum, toBase)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (cv Converter) ConvertFromAnyBaseToDecimal(num string, base int) (int, error) {
	err := validators.CheckBase(base)
	if err != nil {
		return -1, err
	}

	val := 0
	power := 1

	for i := len(num) - 1; i >= 0; i-- {
		digit := cv.charToDecimal(num[i])
		if digit < 0 || digit >= base {
			return -1, errors.New("")
		}
		val += digit * power
		power = power * base
	}

	return val, nil
}

func (cv Converter) charToDecimal(c uint8) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}

	return int(c - 'A' + 10)
}

func (cv Converter) ConvertFromDecimalToBaseX(num int, newBase int) (string, error) {
	err := validators.CheckBase(newBase)
	if err != nil {
		return "", err
	}

	result := strings.Builder{}

	for num > 0 {
		result.WriteString(cv.decimalToChar(num % newBase))
		num /= newBase
	}

	return reverseString(result.String()), nil
}

func (cv Converter) decimalToChar(num int) string {
	if num >= 0 && num <= 9 {
		return string(rune(num + 48))
	}

	return string(rune(num - 10 + 65))
}

func reverseString(s string) string {
	reversed := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		reversed.WriteByte(s[i])
	}

	return reversed.String()
}
