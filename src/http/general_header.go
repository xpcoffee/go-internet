package http

import (
	"fmt"
	"time"
)

const (
	CacheControl     HeaderName = "Cache-Control"
	Connection       HeaderName = "Connection"
	Date             HeaderName = "Date"
	Pragma           HeaderName = "Pragma"
	Trailer          HeaderName = "Trailer"
	TransferEncoding HeaderName = "Transfer-Encoding"
	Upgrade          HeaderName = "Upgrade"
	Via              HeaderName = "Via"
	Warning          HeaderName = "Warning"
)

func (header HeaderName) IsValidGeneralHeader() bool {
	switch header {
	case Date:
		return true
	}
	return false
}

func ParseGeneralHeader(name HeaderName, content string) (Header, error) {

	switch name {
	case Date:
		return (&DateHeader{}).Parse(content)
	}

	return NewUnkownHeader(string(name)), fmt.Errorf("'%s' is not yet an implemented header", name)
}

type CacheControlHeader struct {
	Name    HeaderName
	Content string
}

type ConnectionHeader struct {
	Name    HeaderName
	Content string
}

type DateHeader struct {
	HeaderName HeaderName
	DateTime   time.Time
}

func (header *DateHeader) Parse(content string) (Header, error) {
	datetime, err := time.Parse(time.RFC1123, content)

	if err != nil {
		datetime, err = time.Parse(time.RFC850, content)
	}

	if err != nil {
		datetime, err = time.Parse(time.ANSIC, content)
	}

	if err != nil {
		return header, fmt.Errorf("Could not parse Date header: %s", content)
	}
	header.DateTime = datetime

	return header, nil
}

func (header *DateHeader) Name() HeaderName {
	return header.HeaderName
}

type PragmaHeader struct {
	Name    HeaderName
	Content string
}

type TrailerHeader struct {
	Name    HeaderName
	Content string
}

type TransferEncodingHeader struct {
	Name    HeaderName
	Content string
}

type UpgradeHeader struct {
	Name    HeaderName
	Content string
}

type ViaHeader struct {
	Name    HeaderName
	Content string
}

type WarningHeader struct {
	Name    HeaderName
	Content string
}
