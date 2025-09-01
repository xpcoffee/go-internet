package http

import (
	"fmt"
	"strings"
)

type Header struct {
	Name    HeaderName
	Content string
}

type HeaderName string

const (
	Accept        HeaderName = "Accept"
	Authorization HeaderName = "Authorization"
	ContentType   HeaderName = "Content-Type"
	ContentLength HeaderName = "Content-Length"
)

func (header HeaderName) IsValid() bool {
	switch header {
	case Accept, Authorization, ContentLength, ContentType:
		return true
	}
	return false
}

func ParseHeader(input string) (Header, error) {
	header := Header{}

	splitIndex := strings.Index(strings.Trim(input, " "), " ")
	if splitIndex == -1 {
		return header, fmt.Errorf("Invalid header. Expected 'HeaderName HeaderContent'")
	}

	header.Name = HeaderName(input[:splitIndex])
	header.Content = input[splitIndex+1:]

	if !header.Name.IsValid() {
		return header, fmt.Errorf("'%s' is not a known header name", header.Name)
	}

	return header, nil
}
