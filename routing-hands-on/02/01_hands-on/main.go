package main

import (
	"io"
	"log"
	"net"
)

func respond(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "I see you connected\n")
}



func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go respond(conn)
	}
}