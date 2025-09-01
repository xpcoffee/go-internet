package http

type Request struct {
	RequestLine RequestLine
	Headers     []Header
	MessageBody []byte
}

type URI string
