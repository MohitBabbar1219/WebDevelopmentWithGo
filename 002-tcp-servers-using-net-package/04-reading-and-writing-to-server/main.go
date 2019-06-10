package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
	scn := bufio.NewScanner(conn)
	for scn.Scan() {
		ln := scn.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()
}
