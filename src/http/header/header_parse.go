package header

import (
	"fmt"
	g "internet-protocols/http/header/general"
	m "internet-protocols/http/header/model"
	r "internet-protocols/http/header/request"
	"strings"
)

type Header = m.Header
type HeaderName = m.HeaderName

func ParseHeader(input string) (Header, error) {
	splitIndex := strings.Index(strings.Trim(input, " "), " ")
	if splitIndex == -1 || input[splitIndex-1] != byte(':') {
		return m.NewUnkownHeader("-", "-"), fmt.Errorf("Invalid header. Expected 'HeaderName: HeaderContent'. Got '%s'", input)
	}

	name := HeaderName(input[:splitIndex-1])
	content := input[splitIndex+1:]

	if g.IsValidGeneralHeader(name) {
		return g.ParseGeneralHeader(name, content)
	}

	if r.IsValidRequestHeader(name) {
		return r.ParseRequestHeader(name, content)
	}

	return m.NewUnkownHeader(name, content), nil
}
