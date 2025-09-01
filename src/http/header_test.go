package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseHeader(t *testing.T) {
	result, error := ParseHeader("Accept *")
	expected := Header{Name: Accept, Content: "*"}
	assert.Nil(t, error)
	assert.Equal(t, expected, result)

	result, error = ParseHeader("Content-Length 8")
	expected = Header{Name: ContentLength, Content: "8"}
	assert.Nil(t, error)
	assert.Equal(t, expected, result)

	result, error = ParseHeader("Content-Type application/json")
	expected = Header{Name: ContentType, Content: "application/json"}
	assert.Nil(t, error)
	assert.Equal(t, expected, result)

	result, error = ParseHeader("Authorization Bearer: foobar1234")
	expected = Header{Name: Authorization, Content: "Bearer: foobar1234"}
	assert.Nil(t, error)
	assert.Equal(t, expected, result)

	result, error = ParseHeader("BAD Some content")
	assert.NotNil(t, error)
	assert.Equal(t, "'BAD' is not a known header name", error.Error())

	_, error = ParseHeader("")
	assert.NotNil(t, error)
	assert.Equal(t, "Invalid header. Expected 'HeaderName HeaderContent'", error.Error())

}
