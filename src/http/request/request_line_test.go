package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRequestLine(t *testing.T) {
	result, error := ParseRequestLine("GET https://example.com HTTP/1.1")
	assert.Nil(t, error)

	expected := RequestLine{
		Method:     Get,
		RequestURI: "https://example.com",
		Version:    OneOne,
	}
	assert.Equal(t, expected, result)

	_, error = ParseRequestLine("BAD https://example.com HTTP/1.1")
	assert.NotNil(t, error)

	assert.Equal(t, "'BAD' is not a valid HTTP method", error.Error())
	_, error = ParseRequestLine("GET https://example.com HTTP/8.0")
	assert.NotNil(t, error)
	assert.Equal(t, "'HTTP/8.0' is not a valid HTTP version", error.Error())

	_, error = ParseRequestLine("GET https://example.com BAD")
	assert.NotNil(t, error)
	assert.Equal(t, "'BAD' is not a valid HTTP version", error.Error())

	_, error = ParseRequestLine("")
	assert.NotNil(t, error)
	assert.Equal(t, "Invalid request-line. Expected 'HttpMethod RequestURI HttpVersion'", error.Error())

	_, error = ParseRequestLine("  ")
	assert.NotNil(t, error)
	assert.Equal(t, "Invalid request-line. Expected 'HttpMethod RequestURI HttpVersion'", error.Error())

	_, error = ParseRequestLine("GET  HTTP/1.1")
	assert.NotNil(t, error)
	assert.Equal(t, "The request URI must be defined", error.Error())
}
