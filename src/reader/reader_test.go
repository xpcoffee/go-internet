package reader

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestReadAllCRLF(t *testing.T) {
	multiline_string := "foo\r\nbar\r\nbaz"
	readCloser := io.NopCloser(strings.NewReader(multiline_string))
	br := NewBufferedReader(readCloser)
	result := make([]string, 0)

	out := br.ReadAllCRLF()
	for line := range out {
		result = append(result, line)
	}

	assert.Equal(t, []string{"foo", "bar", "baz"}, result)
}

func TestReadCRLF(t *testing.T) {
	multiline_string := "foo\r\nbarbaz\r\ngrault"
	readCloser := io.NopCloser(strings.NewReader(multiline_string))
	br := NewBufferedReader(readCloser)

	out := make(chan string, 1)
	defer close(out)

	result, has_more := br.ReadCRLF()
	assert.True(t, has_more)
	assert.Equal(t, "foo", result)

	result, has_more = br.ReadCRLF()
	assert.True(t, has_more)
	assert.Equal(t, "barbaz", result)

	result, has_more = br.ReadCRLF()
	assert.False(t, has_more)
	assert.Equal(t, "grault", result)

	result, has_more = br.ReadCRLF()
	assert.False(t, has_more)
	assert.Equal(t, "", result)
}

func TestIterativeRead(t *testing.T) {
	multiline_string := "foo\r\nbar\r\nbaz"
	readCloser := io.NopCloser(strings.NewReader(multiline_string))
	br := NewBufferedReader(readCloser)

	result := make([]string, 0)

	line, has_more := br.ReadCRLF()
	result = append(result, line)
	assert.True(t, has_more)
	assert.Equal(t, "foo", result[0])

	for line := range br.ReadAllCRLF() {
		result = append(result, line)
	}

	assert.Equal(t, []string{"foo", "bar", "baz"}, result)
}

func TestReadAllAsByte(t *testing.T) {
	multiline_string := "foo\rbar\r\nbaz"
	readCloser := io.NopCloser(strings.NewReader(multiline_string))
	br := NewBufferedReader(readCloser)
	result := make([]string, 0)

	out := br.ReadAllAsByte()
	for chunk := range out {
		result = append(result, string(chunk))
	}

	assert.Equal(t, []string{"foo\rbar\r", "\nbaz"}, result)
}
