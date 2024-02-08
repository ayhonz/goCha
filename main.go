package main

import (
	"log"
	"net"
)

func handleConnection(c net.Conn) {
	log.Printf("connection: %s", c.RemoteAddr().String())
	s := "Hello There"

	c.Write([]byte(s))

	for {
		buff := make([]byte, 513)
		_, err := c.Read(buff)
		if err != nil {
			log.Fatalf("Failed to read from conn: %v", err)
		}
		log.Printf("Reading: %s", buff)
	}
}

func main() {
	nl, err := net.Listen("tcp", ":6969")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := nl.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}

}
