package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// 1. listen for proxy request;
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Unable to bind on port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// 2 . handle the request.
		go handleConnection(conn)
	}
}

func handleConnection(src net.Conn) {
	// 1. Connect to target server
	dst, err := net.Dial("tcp", "postman-echo.com:443")
	if err != nil {
		log.Fatalln("Unable to connect to target server")
	}

	defer dst.Close()
	// 2  Copy the src output to the destination conn
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	// 3. Copy the respoonse from the the target server
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
