package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)

		ch <- "hello from goroutine"
		fmt.Println("sent message")
	}()

	msg := <-ch
	fmt.Println("received :", msg)
}
