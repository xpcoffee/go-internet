package model

import (
	"fmt"
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
