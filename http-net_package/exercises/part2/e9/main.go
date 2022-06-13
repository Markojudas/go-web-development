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

	// responding GET / ; GET /apply ; POST /apply
	switch {
	case reqMETHOD == "GET" && reqURL == "/":
		handleIndex(conn)
	case reqMETHOD == "GET" && reqURL == "/apply":
		handleApply(conn)
	case reqMETHOD == "POST" && reqURL == "/apply":
		handleApplyPost(conn)
	default:
		handleDefault(conn)
	}
}

func handleApplyPost(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>POST APPLY</title>
	</head>
	<body>
		<h1>POST APPLY</h1>
		<a href="/">index</a>
		<br>
		<a href="/apply">apply</a>
		<br>
	</body>
	</html>	
	`

	writeReqResp(conn, body)
}

func handleApply(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>GET DOG</title>
	</head>
	<body>
		<h1>GET APPLY</h1>
		<a href="/">index</a>
		<br>
		<a href="/apply">apply</a>
		<br>
		<form action="/apply" method="POST">
		<input type="hidden" value="In my good death">
		<input type="submit" value="submit">
		</form>
	</body>
	</html>	
	`

	writeReqResp(conn, body)
}

func handleIndex(conn net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>GET INDEX</title>
	</head>
	<body>
		<h1>GET INDEX</h1>
		<a href="/">index</a>
		<br>
		<a href="/apply">apply</a>
		<br>
	</body>
	</html>	
	`
	writeReqResp(conn, body)
}

func handleDefault(conn net.Conn) {

	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>EXERCISE 8</title>
	</head>
	<body>
		<h1>Holy Cow I can print want I want really</h1>
	</body>
	</html>	
	`
	writeReqResp(conn, body)
}

func writeReqResp(conn net.Conn, body string) {

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
