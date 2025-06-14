package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := make(chan float64)

	// One sender goroutine
	go sendData(num)

	// One receiver goroutine
	go func() {
		received := receiveData(num)
		fmt.Println("Received:", received)
	}()

	// Wait for goroutines to complete (simple way)
	var input string
	fmt.Scanln(&input)
}

func sendData(out chan<- float64) {
	randmval := rand.Float64()
	fmt.Println("Sending:", randmval)
	out <- randmval
}

func receiveData(in <-chan float64) float64 {
	received := <-in
	return received
}
