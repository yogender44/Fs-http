package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	// to prevent server exiting once it has processed a connection
	for {
		conn, err := listener.Accept() // blocking
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Client connected ", conn)

		go do(conn) // creates a go routine to handle conn processing on a separate thread.
	}
}

func do(conn net.Conn) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf) // blocking
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Processing the request")
	time.Sleep(8 * time.Second)

	conn.Write([]byte("HTTPS/1.1 200 OK\r\nHello World!\r\n")) // blocking system call
	conn.Close()
  // TCP backlog queue configuration
}
