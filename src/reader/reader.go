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

func (br *BufferedReader) ReadAllLines() <-chan string {
	out := make(chan string, 1)

	go func() {
		defer br.Reader.Close()
		defer close(out)

		for {
			line, has_more := br.ReadLine()
			out <- line
			if !has_more {
				break
			}
		}
	}()

	return out
}

// Reads up to the next newline character
// Buffers any extra-read characters
func (br *BufferedReader) ReadLine() (string, bool) {
	line := ""

	if br.Buffer != nil {
		eol := bytes.IndexByte(br.Buffer, '\n')

		if eol != -1 {
			line += string(br.Buffer[:eol])
			br.Buffer = br.Buffer[eol+1:]
			fmt.Println("A")
			return line, true
		}

		line += string(br.Buffer)
	}

	for {
		data := make([]byte, 8)
		n, err := br.Reader.Read(data)

		if errors.Is(err, io.EOF) {
			br.Buffer = nil
			fmt.Println("B")
			return line, false
		}

		eol := bytes.IndexByte(data[:n], '\n')
		if eol != -1 {
			line += string(data[:eol])
			br.Buffer = data[eol+1 : n]
			break
		} else {
			line += string(data[:n])
			br.Buffer = nil
		}

		if errors.Is(err, io.ErrUnexpectedEOF) {
			return line, false
		}
	}

	return line, true
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
	data := make([]byte, 8)
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
