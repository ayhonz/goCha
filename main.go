package main

import (
	"log"
	"net"
)

type MsgType int

const (
	ClientConnected MsgType = iota
	NewMsg
)

type msg struct {
	Text string
	Conn net.Conn
	Type MsgType
}

func serve(messages <-chan msg) {
	var connections []net.Conn

	for {
		msg := <-messages
		switch msg.Type {
		case ClientConnected:
			connections = append(connections, msg.Conn)
		case NewMsg:
			for _, conn := range connections {
				conn.Write([]byte(msg.Text))
			}

		}
	}

}

func handleConnection(c net.Conn, messages chan<- msg) {
	log.Printf("connection: %s", c.RemoteAddr().String())
	s := "Welcome to Hell!!\n"
	c.Write([]byte(s))

	for {
		buf := make([]byte, 513)
		_, err := c.Read(buf)
		if err != nil {
			log.Fatalf("Failed to read from conn: %v", err)
		}
		messages <- msg{Type: NewMsg, Conn: c, Text: string(buf)}
		log.Printf("Reading: %s", buf)
	}
}

func main() {
	nl, err := net.Listen("tcp", ":6969")
	if err != nil {
		panic(err)
	}
	messages := make(chan msg)
	go serve(messages)

	for {
		conn, err := nl.Accept()
		if err != nil {
			panic(err)
		}

		messages <- msg{Conn: conn, Type: ClientConnected}
		go handleConnection(conn, messages)
	}

}
