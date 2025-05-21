package main

import (
	"fmt"
	"time"
)

func main() {

	go sayHello()

	fmt.Println("main function is running")

	time.Sleep(2 * time.Second)
}

func sayHello() {
	fmt.Println("hello from goroutine!")
}
