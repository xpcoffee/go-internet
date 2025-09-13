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
		return (&DateHeader{HeaderName: name}).Parse(content)
	}

	return NewUnkownHeader(name, content), fmt.Errorf("'%s' is not yet an implemented header", name)
}

type DateHeader struct {
	HeaderName HeaderName
	DateTime   time.Time
}

func NewDateHeader(datetime time.Time) *DateHeader {
	return &DateHeader{
		HeaderName: Date,
		DateTime:   datetime,
	}
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

func (header *DateHeader) String() string {
	return fmt.Sprintf("%s: %s", header.HeaderName, header.DateTime.String())
}
