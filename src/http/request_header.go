package http

import (
	"fmt"
	"strconv"
)

const (
	ContentLength HeaderName = "Content-Length"
	Host          HeaderName = "Host"
)

func (header HeaderName) IsValidRequestHeader() bool {
	switch header {
	case ContentLength, Host:
		return true
	}
	return false
}

func ParseRequestHeader(name HeaderName, content string) (Header, error) {

	switch name {
	case ContentLength:
		return (&ContentLengthHeader{HeaderName: name}).Parse(content)
	case Host:
		return (&HostHeader{HeaderName: name}).Parse(content)
	}

	return NewUnkownHeader(name, content), fmt.Errorf("'%s' is not yet an implemented header", name)
}

type ContentLengthHeader struct {
	HeaderName HeaderName
	Value      int
}

func NewContentLengthHeader(value int) *ContentLengthHeader {
	return &ContentLengthHeader{
		HeaderName: ContentLength,
		Value:      value,
	}
}

func (header *ContentLengthHeader) Parse(content string) (Header, error) {
	body_bytes, err := strconv.Atoi(content)
	if err != nil {
		return header, fmt.Errorf("Content-Length value could not be extracted")
	}
	header.Value = body_bytes
	return header, nil
}

func (header *ContentLengthHeader) Name() HeaderName {
	return header.HeaderName
}

func (header *ContentLengthHeader) String() string {
	return fmt.Sprintf("%s: %d", header.Name(), header.Value)
}

type HostHeader struct {
	HeaderName HeaderName
	Value      string
}

func NewHostHeader(value string) *HostHeader {
	return &HostHeader{
		HeaderName: Host,
		Value:      value,
	}
}

func (header *HostHeader) Parse(content string) (Header, error) {
	header.Value = content
	return header, nil
}

func (header *HostHeader) Name() HeaderName {
	return header.HeaderName
}

func (header *HostHeader) String() string {
	return fmt.Sprintf("%s: %s", header.Name(), header.Value)
}
