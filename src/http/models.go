package http

type Request struct {
	RequestLine RequestLine
	Headers     []Header
	MessageBody []byte
}

type URI string

type HttpMethod string

const (
	Get    HttpMethod = "GET"
	Post   HttpMethod = "POST"
	Put    HttpMethod = "PUT"
	Delete HttpMethod = "Delete"
	Option HttpMethod = "Option"
)

func (m HttpMethod) IsValid() bool {
	switch m {
	case Get, Post, Put, Delete, Option:
		return true
	}
	return false
}

type HttpVersion string

const (
	OneOne HttpVersion = "HTTP/1.1"
)

func (m HttpVersion) IsValid() bool {
	switch m {
	case OneOne:
		return true
	}
	return false
}
