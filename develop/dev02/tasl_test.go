package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
			err:      false,
		},
		{
			input:    "abcd",
			expected: "abcd",
			err:      false,
		},
		{
			input:    "45",
			expected: "",
			err:      true,
		},
		{
			input:    "",
			expected: "",
			err:      false,
		},
		{
			input:    "qwe\\4\\5",
			expected: "qwe45",
			err:      false,
		},
		{
			input:    "qwe\\45",
			expected: "qwe44444",
			err:      false,
		},
		{
			input:    "qwe\\\\5",
			expected: "qwe\\\\\\\\\\",
			err:      false,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			out, err := Unpack(&test.input)
			if test.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, test.expected, out)
		})
	}
}
