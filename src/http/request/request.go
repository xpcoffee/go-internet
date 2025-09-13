package http

import (
	"fmt"
	"internet-protocols/reader"

	h "internet-protocols/http/header"
	m "internet-protocols/http/header/model"
	r "internet-protocols/http/header/request"
)

type Request struct {
	RequestLine RequestLine
	Headers     map[m.HeaderName]m.Header
	MessageBody string
}

func (request Request) String() string {
	return fmt.Sprintf("%s %+v %s", request.RequestLine, request.Headers, request.MessageBody)
}

type URI string

func ParseRequest(br *reader.BufferedReader) (Request, error) {
	request := Request{}

	requestLineStr, has_more := br.ReadCRLF()
	requestLine, err := ParseRequestLine(requestLineStr)
	request.RequestLine = requestLine
	if err != nil {
		return request, err
	}
	if !has_more {
		return request, fmt.Errorf("Incomplete request")
	}

	request.Headers = make(map[m.HeaderName]m.Header)
	for {
		headerStr, has_more := br.ReadCRLF()
		if headerStr == "" || !has_more {
			break
		}

		header, error := h.ParseHeader(headerStr)
		if error != nil {
			return request, error
		}
		request.Headers[header.Name()] = header
	}

	content_length, ok := request.Headers[r.ContentLength].(*r.ContentLengthHeader)
	if ok && content_length != nil {
		buffer_bytes := len(br.Buffer)
		bytes_to_read := *&content_length.Value - buffer_bytes
		data := make([]byte, bytes_to_read)
		n, err := br.Reader.Read(data)
		if err != nil {
			return request, fmt.Errorf("Could not parse body: %s", err.Error())
		}

		if n != bytes_to_read {
			return request, fmt.Errorf("Body does not contain data the size of Content-Length")
		}

		body := append(br.Buffer, data...)
		request.MessageBody = string(body)
	}

	return request, nil
}
