package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func readLinesFromReader(reader io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer close(out)
		defer reader.Close()

		line := ""
		for {
			data := make([]byte, 8)
			n, err := reader.Read(data)
			if errors.Is(err, io.EOF) {
				if len(line) > 0 {
					out <- line
				}
				break
			}

			eol := bytes.IndexByte(data[:n], '\n')
			if eol != -1 {
				line += string(data[:eol])
				out <- line
				line = string(data[eol+1 : n])
			} else {
				line += string(data[:n])
			}

			if errors.Is(err, io.ErrUnexpectedEOF) {
				if len(line) > 0 {
					out <- line
				}
				break
			}
		}

	}()

	return out
}

func main() {
	f, err := os.Open("data/full-request.txt")
	if err != nil {
		panic(err)
	}

	for line := range readLinesFromReader(f) {
		fmt.Println(line)
	}
}
