package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	name string
	ch   chan<- string
}

var (
	messages = make(chan string)
	entering = make(chan client)
	leaving  = make(chan client)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Chat server started on localhost:8000")

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	who := conn.RemoteAddr().String()
	client := client{name: who, ch: ch}

	go clientWriter(conn, ch)

	ch <- "You are " + who
	entering <- client
	messages <- who + " has arrived"

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- client
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			// ✅ Show the list of current clients to the new client
			cli.ch <- "Currently in chat:"
			for c := range clients {
				cli.ch <- " - " + c.name
			}
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}
