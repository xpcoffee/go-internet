package request

import (
	"fmt"
	m "internet-protocols/http/header/model"
)

const Host m.HeaderName = "Host"

type HostHeader struct {
	HeaderName m.HeaderName
	Value      string
}

func NewHostHeader(value string) *HostHeader {
	return &HostHeader{
		HeaderName: Host,
		Value:      value,
	}
}

func (header *HostHeader) Parse(content string) (m.Header, error) {
	header.Value = content
	return header, nil
}

func (header *HostHeader) Name() m.HeaderName {
	return header.HeaderName
}

func (header *HostHeader) String() string {
	return fmt.Sprintf("%s: %s", header.Name(), header.Value)
}
