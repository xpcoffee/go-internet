package http

import (
	"fmt"
	"strings"
)

func ParseRequestLine(input string) (RequestLine, error) {
	requestLine := RequestLine{}

	parts := strings.Split(strings.Trim(input, " "), " ")
	if len(parts) != 3 {
		return requestLine, fmt.Errorf("Invalid request-line. Expected 'HttpMethod RequestURI HttpVersion'")
	}

	method := HttpMethod(parts[0])
	requestLine.Method = method
	if !requestLine.Method.IsValid() {
		return requestLine, fmt.Errorf("'%s' is not a valid HTTP method", requestLine.Method)
	}

	requestLine.RequestURI = URI(parts[1])
	if len(requestLine.RequestURI) == 0 {
		return requestLine, fmt.Errorf("The request URI must be defined")
	}

	requestLine.Version = HttpVersion(parts[2])
	if !requestLine.Version.IsValid() {
		return requestLine, fmt.Errorf("'%s' is not a valid HTTP version", requestLine.Version)
	}

	return requestLine, nil
}
