package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		byteSlice := []byte(ln)
		r := rot13(byteSlice)

		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
	}
}

// rot13 gets a byte slice "jumbles" the characters of a string
// if the ascii of a character is less or equal to 109 we add 13
// if the ascii of a character is greater than 109 we subtract 13
//eg : n - ascii 110, 110  - 13 = 97 which is a
//eg : m - ascii 109, 109 + 13 = 122 which is z
func rot13(bs []byte) []byte {

	var r13 = make([]byte, len(bs)) //making a new byte slice

	for idx, val := range bs {
		// ascii 97 - 122 - lowe case characters,
		if val <= 109 {
			r13[idx] = val + 13
		} else {
			r13[idx] = val - 13
		}
	}

	return r13
}
