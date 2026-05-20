package code

import (
	errs "code/internal/errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathSize_File(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		name      string
		target    string
		recursive bool
		human     bool
		all       bool
		expected  string
	}{
		{name: "unhuman", target: "testdata/file49kb", expected: fmt.Sprintf("%dB", 49*1024)},
		{name: "human", target: "testdata/file49kb", human: true, expected: "49.0KB"},
		{name: "hidden file", target: "testdata/.file25kb", expected: fmt.Sprintf("%dB", 25*1024)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetPathSize(tt.target, tt.recursive, tt.human, tt.all)
			a.Nil(err)
			a.Equal(tt.expected, result)
		})
	}
}

func TestGetPathSize_Directory(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		name      string
		target    string
		recursive bool
		human     bool
		all       bool
		expected  string
	}{
		{name: "unhuman", target: "testdata", expected: fmt.Sprintf("%dB", 100*1024)},
		{name: "human", target: "testdata", human: true, expected: "100.0KB"},
		{name: "with hidden files", target: "testdata", all: true, expected: fmt.Sprintf("%dB", 125*1024)},
		{name: "recursive", target: "testdata", recursive: true, expected: fmt.Sprintf("%dB", 200*1024)},
		{name: "with recursive and hidden files", target: "testdata", recursive: true, all: true, expected: fmt.Sprintf("%dB", 225*1024)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetPathSize(tt.target, tt.recursive, tt.human, tt.all)
			a.Nil(err)
			a.Equal(tt.expected, result)
		})
	}
}

func TestGetPathSize_FileNotExist(t *testing.T) {
	a := assert.New(t)

	result, err := GetPathSize("testdata/nonexistent", false, false, false)
	a.Error(err)
	a.Empty(result)
}

func TestGetPathSize_EmptyPath(t *testing.T) {
	a := assert.New(t)

	result, err := GetPathSize("", false, false, false)
	a.ErrorIs(err, errs.ErrEmptyPath)
	a.Empty(result)
}
