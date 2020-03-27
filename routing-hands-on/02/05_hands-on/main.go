package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)


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

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			//when ln is empty the headers is done
			if ln == "" {
				fmt.Println("This is the end of req headers")
				break
			}
			fmt.Println(ln)
			
		}
		fmt.Println("Code got here?")
		io.WriteString(conn, "I see you connected\n")
		
		conn.Close()
	}
}