package request

import (
	"fmt"
	m "internet-protocols/http/header/model"
)

func IsValidRequestHeader(header m.HeaderName) bool {
	switch header {
	case ContentLength, Host:
		return true
	}
	return false
}

func ParseRequestHeader(name m.HeaderName, content string) (m.Header, error) {

	switch name {
	case ContentLength:
		return (&ContentLengthHeader{HeaderName: name}).Parse(content)
	case Host:
		return (&HostHeader{HeaderName: name}).Parse(content)
	}

	return m.NewUnkownHeader(name, content), fmt.Errorf("'%s' is not yet an implemented header", name)
}
