package http

import (
	"errors"
	"fmt"
	"internet-protocols/reader"
	"io"
	"strconv"
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
	requestLine, err := ParseRequestLine(requestLineStr)
	request.RequestLine = requestLine
	if err != nil {
		return request, err
	}
	if !has_more {
		return request, fmt.Errorf("Incomplete request")
	}

	body_bytes := 0

	request.Headers = make([]Header, 0)
	for {
		headerStr, has_more := br.ReadLine()
		if headerStr == "" || !has_more {
			break
		}

		header, error := ParseHeader(headerStr)
		request.Headers = append(request.Headers, header)
		if error != nil {
			return request, error
		}

		if header.Name == ContentLength {
			body_bytes, err = strconv.Atoi(header.Content)
			if err != nil {
				return request, fmt.Errorf("Content-Length value could not be extracted")
			}
		}
	}

	if body_bytes > 0 {
		body := make([]byte, body_bytes)
		n, err := br.Reader.Read(body)
		if err != nil {
			return request, fmt.Errorf("Could not parse body")
		}

		overflow := make([]byte, 1)
		m, err := br.Reader.Read(overflow)
		if errors.Is(err, io.EOF) {
			return request, fmt.Errorf("Could not parse body")
		}

		if n != body_bytes || m > 0 {
			return request, fmt.Errorf("Body content is not the size of Content-Length")
		}

		request.MessageBody = string(body)
	}

	return request, nil
}
