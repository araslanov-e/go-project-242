package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		target   string
		expected string
	}{
		{"testdata/file49kb", "49.0KB"},
		{"testdata", "100.0KB"},
	}

	for _, test := range tests {
		result, err := GetPathSize(test.target)
		a.Nil(err)
		a.Equal(test.expected, result)
	}

	result, err := GetPathSize("")
	a.EqualError(err, "path is empty")
	a.Empty(result)
}
