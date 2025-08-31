package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("data/full-request.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	str := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if errors.Is(err, io.EOF) {
			break
		}

		str += string(data[:n])
		if errors.Is(err, io.ErrUnexpectedEOF) {
			break
		}
		fmt.Printf("read: %s\n", string(data[:n]))
	}

	fmt.Println(str)
}
