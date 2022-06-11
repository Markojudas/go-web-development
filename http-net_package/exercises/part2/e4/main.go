package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	//Listener
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	//infinite loop that accepts connections!
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

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			fmt.Println("We have reached the end!")
			break
		}
	}
	defer conn.Close()

	fmt.Println("Code Got Here")
	io.WriteString(conn, "I see you connected")
}
