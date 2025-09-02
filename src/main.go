package main

import (
	"fmt"
	"internet-protocols/http"
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

		request, err := http.ParseRequest(reader.NewBufferedReader(conn))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(request)
		}

		var response = []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello World!")
		conn.Write(response)
		conn.Close()
	}
}
