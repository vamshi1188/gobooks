package main

import (
	"fmt"
	"time"
)

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	for {
		printer := <-squares
		fmt.Println(printer)
		T := 1 * time.Millisecond
		time.Sleep(T)
		if printer == 9597604 {
			break
		}

	}
	close(naturals)
	close(squares)
}
