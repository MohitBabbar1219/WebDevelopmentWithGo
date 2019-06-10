package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	request(conn)
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scn := bufio.NewScanner(conn)
	for scn.Scan() {
		ln := scn.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			fmt.Println("***METHOD", m)
		}
		if ln == "" {
			break
		}
		i += 1
	}
}

func response(conn net.Conn) {
	body := `<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Go Template</title>
</head>
<body>
    <h1>Basic</h1>
</body>
</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}