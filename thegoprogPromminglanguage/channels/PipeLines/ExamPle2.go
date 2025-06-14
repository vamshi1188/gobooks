package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	person1msg := make(chan string)
	person2msg := make(chan string)

	fmt.Print("hello type person1 msg: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()

	// Send message from person1 to person2
	go func() {
		person1msg <- message
	}()

	// Receive from person1 and send to person2
	go func() {
		received := <-person1msg
		person2msg <- received
	}()

	// Receive and print person2's message
	fmt.Println("Message received by person2:", <-person2msg)
}
