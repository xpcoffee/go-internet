package reader

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestReadAllAsString(t *testing.T) {
	multiline_string := "foo\nbar\nbaz"
	readCloser := io.NopCloser(strings.NewReader(multiline_string))
	br := NewBufferedReader(readCloser)
	result := make([]string, 0)

	out := br.ReadAllLines()
	for line := range out {
		result = append(result, line)
	}

	assert.Equal(t, []string{"foo", "bar", "baz"}, result)
}

func TestReadLine(t *testing.T) {
	multiline_string := "foo\nbar\nbaz"
	readCloser := io.NopCloser(strings.NewReader(multiline_string))
	br := NewBufferedReader(readCloser)

	out := make(chan string, 1)
	defer close(out)

	result, has_more := br.ReadLine()
	assert.True(t, has_more)
	assert.Equal(t, "foo", result)

	result, has_more = br.ReadLine()
	assert.True(t, has_more)
	assert.Equal(t, "bar", result)

	result, has_more = br.ReadLine()
	assert.False(t, has_more)
	assert.Equal(t, "baz", result)

	result, has_more = br.ReadLine()
	assert.False(t, has_more)
	assert.Equal(t, "", result)
}

func TestIterativeRead(t *testing.T) {
	multiline_string := "foo\nbar\nbaz"
	readCloser := io.NopCloser(strings.NewReader(multiline_string))
	br := NewBufferedReader(readCloser)

	result := make([]string, 0)

	line, has_more := br.ReadLine()
	result = append(result, line)
	assert.True(t, has_more)
	assert.Equal(t, "foo", result[0])

	for line := range br.ReadAllLines() {
		result = append(result, line)
	}

	assert.Equal(t, []string{"foo", "bar", "baz"}, result)
}

func TestReadAllAsByte(t *testing.T) {
	multiline_string := "foo\nbar\nbaz"
	readCloser := io.NopCloser(strings.NewReader(multiline_string))
	br := NewBufferedReader(readCloser)
	result := make([]string, 0)

	out := br.ReadAllAsByte()
	for chunk := range out {
		result = append(result, string(chunk))
	}

	assert.Equal(t, []string{"foo\nbar\n", "baz"}, result)
}
