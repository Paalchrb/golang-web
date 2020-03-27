package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	defer conn.Close()
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			//when ln is empty the headers is done
			fmt.Println("This is the end of req headers")
			break
		}
		fmt.Println(ln)
	}
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
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
		go serve(conn)
	}
}