package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils_StringToInt(t *testing.T) {
	t.Parallel()

	var (
		tests = []struct {
			name      string
			input     string
			expected  int
			assertion assert.ComparisonAssertionFunc
		}{
			{"convert integer number", "123", 123, assert.Equal},
			{"convert float number", "456.78", 456, assert.Equal},
			{"convert non-empty string", "abc", 0, assert.Equal},
			{"convert empty string", "", 0, assert.Equal},
		}
		utils = new(utils)
	)

	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			result, _ := utils.StringToInt(ts.input)
			ts.assertion(t, ts.expected, result)
		})
	}
}
