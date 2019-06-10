package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		_, _ = io.WriteString(conn, "\nHello from TCP server\n")
		_, _ = fmt.Fprintln(conn, "How is your day?")
		_, _ = fmt.Fprintf(conn, "%v", "Well, I hope.")
		_ = conn.Close()
	}
}
