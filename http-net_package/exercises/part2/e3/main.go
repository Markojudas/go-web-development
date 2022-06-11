package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	//listener
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	//infinite loop!
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)

			//if ln is an empty string
			if ln == "" {
				fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
				break
			}
		}

		fmt.Println("Code Got Here // when I stopped the browser")
		io.WriteString(conn, "I see you connected")

		conn.Close()
	}
}
