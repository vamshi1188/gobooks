package main

import (
	"bufio"
	"fmt"
	// "go/scanner"
	"log"
	"net"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8001")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go handleEcho(conn)
	}
}

func handleEcho(c net.Conn) {

	defer c.Close()
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		input := scanner.Text()

		fmt.Println(c, strings.ToUpper(input))
	}
}
