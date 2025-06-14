package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 69
	}()
	x := <-ch

	fmt.Println("Received: ", x)

}
