package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	//create a listener
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	//infinite loop for connection
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	//create a scanner
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			fmt.Println("End of line!")
			break
		}
	}
	defer conn.Close()

	body := "Check ot the response body payload"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")            //status line
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) //resonse headers
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")      //response headers
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
