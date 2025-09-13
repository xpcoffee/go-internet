package http

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseHeader(t *testing.T) {
	result, error := ParseHeader("Content-Length: 8")
	assert.Nil(t, error)
	assert.Equal(t, NewContentLengthHeader(8), result)

	result, error = ParseHeader("Date: Tue, 29 Oct 2024 16:56:32 GMT")
	assert.Nil(t, error)
	assert.Equal(t, time.Date(2024, 10, 29, 16, 56, 32, 0, time.UTC), result.(*DateHeader).DateTime.UTC())

	_, error = ParseHeader("")
	assert.NotNil(t, error)
	assert.Equal(t, "Invalid header. Expected 'HeaderName: HeaderContent'. Got ''", error.Error())

}
