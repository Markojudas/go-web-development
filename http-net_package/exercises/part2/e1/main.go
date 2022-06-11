package main

import (
	"io"
	"log"
	"net"
)

func main() {
	//creeating a listener
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	//infinite loop to maintain the connection open

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		io.WriteString(conn, "I see you've connected!\n")
		conn.Close()
	}
}
