package main

import "fmt"

type counter struct {
	count int
}

func (c counter) IncrementByValue() {
	c.count++
}

func (c *counter) IncrementByPointer() {
	c.count++
}

func main() {
	counter := counter{count: 0}

	counter.IncrementByValue()
	fmt.Println("after value method", counter.count)

	counter.IncrementByPointer()
	fmt.Println("after pointer method", counter.count)

	counter.IncrementByPointer()
	fmt.Println("after pointer method", counter.count)
}
