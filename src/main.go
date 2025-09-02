package main

import (
	"fmt"
	"internet-protocols/reader"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		for line := range reader.NewBufferedReader(conn).ReadAllLines() {
			fmt.Println(line)
		}
	}
}
