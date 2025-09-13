package http

import (
	"github.com/stretchr/testify/assert"
	"internet-protocols/reader"
	"io"
	"strings"
	"testing"
)

func TestParseRequest(t *testing.T) {
	request_str :=
		"GET http://www.example.com HTTP/1.1\r\n" +
			"Host: localhost\r\n" +
			"Content-Length: 17\r\n" +
			"\r\n" +
			`{"hello":"world"}`

	expected := Request{
		RequestLine: RequestLine{
			Method:     Get,
			RequestURI: "http://www.example.com",
			Version:    OneOne,
		},
		Headers: map[HeaderName]Header{
			ContentLength: NewContentLengthHeader(17),
			Host:          NewHostHeader("localhost"),
		},
		MessageBody: `{"hello":"world"}`,
	}

	readCloser := io.NopCloser(strings.NewReader(request_str))
	br := reader.NewBufferedReader(readCloser)
	result, err := ParseRequest(br)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
