package code

import (
	errs "code/internal/errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		target    string
		recursive bool
		human     bool
		all       bool
		expected  string
	}{
		{"testdata/file49kb", false, true, false, "49.0KB"},
		{"testdata/file49kb", false, false, false, fmt.Sprintf("%dB", 49*1024)},
		{"testdata", false, true, false, "100.0KB"},
		{"testdata", false, false, false, fmt.Sprintf("%dB", 100*1024)},
		{"testdata", false, false, true, fmt.Sprintf("%dB", 125*1024)},
		{"testdata", true, false, false, fmt.Sprintf("%dB", 200*1024)},
		{"testdata", true, false, true, fmt.Sprintf("%dB", 225*1024)},
	}

	for _, test := range tests {
		result, err := GetPathSize(test.target, test.recursive, test.human, test.all)
		a.Nil(err)
		a.Equal(test.expected, result)
	}
}

func TestGetPathSize_EmptyPath(t *testing.T) {
	a := assert.New(t)

	result, err := GetPathSize("", false, false, false)
	a.ErrorIs(err, errs.ErrEmptyPath)
	a.Empty(result)
}
