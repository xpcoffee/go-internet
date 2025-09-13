package general

import (
	"fmt"
	m "internet-protocols/http/header/model"
)

const ()

func IsValidGeneralHeader(header m.HeaderName) bool {
	switch header {
	case Date:
		return true
	}
	return false
}

func ParseGeneralHeader(name m.HeaderName, content string) (m.Header, error) {

	switch name {
	case Date:
		return (&DateHeader{HeaderName: name}).Parse(content)
	}

	return m.NewUnkownHeader(name, content), fmt.Errorf("'%s' is not yet an implemented header", name)
}
