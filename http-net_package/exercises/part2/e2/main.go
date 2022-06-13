package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	//first create listener
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	//infinte loop for connection
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		//creating a scanner to READ from connection
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected") //this is never reached - caught in the loop!

		conn.Close()

	}
}
