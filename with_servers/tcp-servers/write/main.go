package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	//listen
	li, err := net.Listen("tcp", ":8080") //localhost port 8080

	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept() //waits and accept next connection
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP Server\n") //sending message to the connection
		fmt.Fprintln(conn, "how is your day")
		fmt.Fprintf(conn, "%v", "well, I hope!")
		fmt.Fprintln(conn)

		conn.Close()
	}
}
