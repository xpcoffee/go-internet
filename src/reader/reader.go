package reader

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type BufferedReader struct {
	Reader io.ReadCloser
	// When reading, we may "over-read" past a symbol we care about like a newline.
	// The buffer allows us to hold onto those bytes for the next operation.
	Buffer []byte
}

func NewBufferedReader(reader io.ReadCloser) *BufferedReader {
	return &BufferedReader{
		Reader: reader,
	}
}

func (br *BufferedReader) ReadAllCRLF() <-chan string {
	out := make(chan string, 1)

	go func() {
		defer br.Reader.Close()
		defer close(out)

		for {
			line, has_more := br.ReadCRLF()
			out <- line
			if !has_more {
				break
			}
		}
	}()

	return out
}

const DEFAULT_CHUNK_SIZE_BYTES = 8

// Reads up to the next newline character
// Buffers any extra-read characters
func (br *BufferedReader) ReadCRLF() (string, bool) {
	has_more := true

	for {
		// first check what's in the current buffer
		eol := bytes.IndexByte(br.Buffer, '\n')
		crlf_found := eol > 0 && br.Buffer[eol-1] == byte('\r')
		if crlf_found {
			data := br.Buffer[:eol-1]

			if len(br.Buffer) > eol {
				left_over := br.Buffer[eol+1:]
				br.Buffer = left_over
			}

			return string(data), has_more
		}

		// then read more data
		data := make([]byte, DEFAULT_CHUNK_SIZE_BYTES)
		n, err := br.Reader.Read(data)
		br.Buffer = append(br.Buffer, data[:n]...)

		if err != nil {
			fmt.Printf("error %s\n", err.Error())
		}

		if err != nil || n < len(data) {
			has_more = false
		}

		if !has_more {
			return string(br.Buffer), has_more
		}
	}
}

func (br *BufferedReader) ReadAllAsByte() <-chan []byte {
	out := make(chan []byte, 1)

	go func() {
		defer br.Reader.Close()
		defer close(out)
		for {
			chunk, has_more := br.ReadChunk()
			out <- chunk
			if !has_more {
				break
			}
		}
	}()

	return out
}

// Reads the next chunk of data
// Appends it to the buffer and clears the buffer, if any
func (br *BufferedReader) ReadChunk() ([]byte, bool) {
	data := make([]byte, DEFAULT_CHUNK_SIZE_BYTES)
	n, err := br.Reader.Read(data)
	has_more := n == len(data)
	data = data[:n]

	if br.Buffer != nil {
		data = append(br.Buffer, data...)
		br.Buffer = nil
	}

	if errors.Is(err, io.EOF) {
		has_more = false
	}

	if err != nil {
		fmt.Printf("return %t %s\n", has_more, err.Error())
	}
	return data, has_more
}
