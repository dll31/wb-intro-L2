package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTime(t *testing.T) {
	tests := []struct {
		result string
		server string
		err    bool
	}{
		{
			result: "OK",
			server: "0.beevik-ntp.pool.ntp.org",
			err:    false,
		},
		{
			result: "bad server url",
			server: "Bad url",
			err:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			_, err := getTime(test.server)
			if test.err {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
