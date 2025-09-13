package request

import (
	"fmt"
	m "internet-protocols/http/header/model"
	"strconv"
)

const ContentLength m.HeaderName = "Content-Length"

type ContentLengthHeader struct {
	HeaderName m.HeaderName
	Value      int
}

func NewContentLengthHeader(value int) *ContentLengthHeader {
	return &ContentLengthHeader{
		HeaderName: ContentLength,
		Value:      value,
	}
}

func (header *ContentLengthHeader) Parse(content string) (m.Header, error) {
	body_bytes, err := strconv.Atoi(content)
	if err != nil {
		return header, fmt.Errorf("Content-Length value could not be extracted")
	}
	header.Value = body_bytes
	return header, nil
}

func (header *ContentLengthHeader) Name() m.HeaderName {
	return header.HeaderName
}

func (header *ContentLengthHeader) String() string {
	return fmt.Sprintf("%s: %d", header.Name(), header.Value)
}
