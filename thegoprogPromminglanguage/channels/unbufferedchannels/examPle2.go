package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	tcpConn := conn.(*net.TCPConn)
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, tcpConn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(tcpConn, os.Stdin)
	tcpConn.CloseWrite() // Only close write-half
	<-done               // Wait until the goroutine finishes
}

// mustCopy copies data from src to dst using io.Copy.
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
