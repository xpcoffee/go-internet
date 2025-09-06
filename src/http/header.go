package http

import (
	"fmt"
	"strings"
)

type Header interface {
	Name() HeaderName
	String() string
	Parse(content string) (Header, error)
}

type HeaderName string

func (header HeaderName) String() string {
	return string(header)
}

func ParseHeader(input string) (Header, error) {
	splitIndex := strings.Index(strings.Trim(input, " "), " ")
	if splitIndex == -1 || input[splitIndex-1] != byte(':') {
		return NewUnkownHeader("-", "-"), fmt.Errorf("Invalid header. Expected 'HeaderName: HeaderContent'. Got '%s'", input)
	}

	name := HeaderName(input[:splitIndex-1])
	content := input[splitIndex+1:]

	if name.IsValidGeneralHeader() {
		return ParseGeneralHeader(name, content)
	}

	if name.IsValidRequestHeader() {
		return ParseRequestHeader(name, content)
	}

	return NewUnkownHeader(name, content), nil
}

type UnknownHeader struct {
	HeaderName HeaderName
	Value      string
}

func (header *UnknownHeader) Parse(content string) (Header, error) {
	// noop
	return header, nil
}

func (header *UnknownHeader) String() string {
	return fmt.Sprintf("%s: %s", header.HeaderName, header.Value)
}

func NewUnkownHeader(name HeaderName, content string) *UnknownHeader {
	return &UnknownHeader{HeaderName: name, Value: content}
}

func (header *UnknownHeader) Name() HeaderName {
	return header.HeaderName
}
