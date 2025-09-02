package http

import (
	"fmt"
	"internet-protocols/reader"
)

type Request struct {
	RequestLine RequestLine
	Headers     []Header
	MessageBody string
}

type URI string

func ParseRequest(br *reader.BufferedReader) (Request, error) {
	request := Request{}

	requestLineStr, has_more := br.ReadLine()
	fmt.Printf("hasmor %t - %s\n", has_more, requestLineStr)
	requestLine, err := ParseRequestLine(requestLineStr)
	fmt.Printf("Requestline %s\n", requestLine)
	request.RequestLine = requestLine
	if err != nil {
		return request, err
	}
	if !has_more {
		return request, fmt.Errorf("Incomplete request")
	}

	request.Headers = make([]Header, 0)
	for {
		headerStr, hase_more := br.ReadLine()
		if headerStr == "" {
			break
		}

		header, error := ParseHeader(headerStr)
		request.Headers = append(request.Headers, header)
		if error != nil {
			return request, error
		}

		if !hase_more {
			return request, fmt.Errorf("Expected empty CRLF after headers")
		}
	}

	body := make([]byte, 0)
	out := br.ReadAllAsByte()
	for chunk := range out {
		body = append(body, chunk...)
	}
	if len(body) > 0 {
		request.MessageBody = string(body)
	}

	return request, nil
}
