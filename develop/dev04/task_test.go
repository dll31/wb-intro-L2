package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagramSearch(t *testing.T) {
	OKExpectedOutput := make(map[string][]string)
	OKExpectedOutput["листок"] = append(OKExpectedOutput["листок"], "столик")

	tests := []struct {
		name           string
		input          []string
		expectedOutput map[string][]string
	}{
		{
			name:           "OK",
			input:          []string{"листок", "столик", "столик", "слоник", "дом"},
			expectedOutput: OKExpectedOutput,
		},
		{
			name:           "empty",
			input:          []string{},
			expectedOutput: make(map[string][]string),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			realOutput := AnagramSearch(&test.input)
			assert.Equal(t, test.expectedOutput, realOutput)
		})
	}
}
