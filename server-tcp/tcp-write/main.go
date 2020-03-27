package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err  := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("TCP listening to port 8080")
	defer li.Close()

	for {
		conn, err :=  li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v\n", "Well, I hope!")

		conn.Close()
	}
}