package http

import (
	"fmt"
	"strings"
)

type Header interface {
	Name() HeaderName
	Parse(content string) (Header, error)
}

type HeaderName string

func (header HeaderName) IsValid() bool {
	// Disable for now - too many in spec
	// switch header {
	// case Accept, Authorization, ContentLength, ContentType, Host, Date, UserAgent, ContentEncoding, Pragma, Connection:
	// 	return true
	// }
	return true
}

func ParseHeader(input string) (Header, error) {
	header := &UnknownHeader{Value: input}
	splitIndex := strings.Index(strings.Trim(input, " "), " ")
	if splitIndex == -1 || input[splitIndex-1] != byte(':') {
		return NewUnkownHeader(input), fmt.Errorf("Invalid header. Expected 'HeaderName: HeaderContent'. Got '%s'", input)
	}

	name := HeaderName(input[:splitIndex-1])
	content := input[splitIndex+1:]

	if name.IsValidGeneralHeader() {
		return ParseGeneralHeader(name, content)
	}

	switch name {
	default:
		return header, nil
	}
}

type UnknownHeader struct {
	HeaderName HeaderName
	Value      string
}

func (header *UnknownHeader) Parse(content string) (Header, error) {
	// noop
	return header, nil
}

func NewUnkownHeader(headerStr string) *UnknownHeader {
	return &UnknownHeader{Value: headerStr}
}

func (header *UnknownHeader) Name() HeaderName {
	return header.HeaderName
}
