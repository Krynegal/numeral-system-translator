package converter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertFromAnyBaseToDecimal(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		tests := []struct {
			number   string
			base     int
			expected int
		}{
			{
				number:   "10",
				base:     10,
				expected: 10,
			},
			{
				number:   "1011",
				base:     2,
				expected: 11,
			},
			{
				number:   "2B",
				base:     16,
				expected: 43,
			},
		}

		converter := Converter{}

		for _, tt := range tests {
			result, err := converter.ConvertFromAnyBaseToDecimal(tt.number, tt.base)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		}
	})
}

func TestConvertFromDecimalToBaseX(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		tests := []struct {
			number   int
			toBase   int
			expected string
		}{
			{
				number:   21,
				toBase:   9,
				expected: "23",
			},
			{
				number:   45,
				toBase:   13,
				expected: "36",
			},
			{
				number:   251,
				toBase:   16,
				expected: "FB",
			},
		}

		converter := Converter{}

		for _, tt := range tests {
			result, err := converter.ConvertFromDecimalToBaseX(tt.number, tt.toBase)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		}
	})
}
