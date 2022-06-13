package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	//create a scanner
	scanner := bufio.NewScanner(conn)

	var reqMETHOD, reqURL string

	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			xs := strings.Fields(ln)          //breaking the line into a slice of strings
			reqMETHOD = xs[0]                 //method (GET)
			reqURL = xs[1]                    //URL
			fmt.Println("METHOD:", reqMETHOD) // printing to terminal
			fmt.Println("URL:", reqURL)       // printing to terminal
		}
		if ln == "" {
			fmt.Println("End of line!")
			break
		}
		i++
	}
	defer conn.Close()

	body := "Check ot the response body payload"
	body += "\n"
	body += reqMETHOD
	body += "\n"
	body += reqURL
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
