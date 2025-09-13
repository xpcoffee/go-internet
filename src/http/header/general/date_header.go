package general

import (
	"fmt"
	m "internet-protocols/http/header/model"
	"time"
)

const Date m.HeaderName = "Date"

type DateHeader struct {
	HeaderName m.HeaderName
	DateTime   time.Time
}

func NewDateHeader(datetime time.Time) *DateHeader {
	return &DateHeader{
		HeaderName: Date,
		DateTime:   datetime,
	}
}

func (header *DateHeader) Parse(content string) (m.Header, error) {
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

func (header *DateHeader) Name() m.HeaderName {
	return header.HeaderName
}

func (header *DateHeader) String() string {
	return fmt.Sprintf("%s: %s", header.HeaderName, header.DateTime.String())
}
