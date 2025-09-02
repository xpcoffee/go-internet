package http

import (
	"github.com/stretchr/testify/assert"
	"internet-protocols/reader"
	"io"
	"strings"
	"testing"
)

func TestParseRequest(t *testing.T) {
	request_str := `GET http://www.example.com HTTP/1.1
Host: localhost
Content-Length: 17

{"hello":"world"}`

	expected := Request{
		RequestLine: RequestLine{
			Method:     Get,
			RequestURI: "http://www.example.com",
			Version:    OneOne,
		},
		Headers: []Header{
			{Name: Host, Content: "localhost"},
			{Name: ContentLength, Content: "17"},
		},
		MessageBody: `{"hello":"world"}`,
	}

	readCloser := io.NopCloser(strings.NewReader(request_str))
	br := reader.NewBufferedReader(readCloser)
	result, err := ParseRequest(br)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
