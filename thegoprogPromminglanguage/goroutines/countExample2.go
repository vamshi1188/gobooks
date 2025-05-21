package main

import (
	"fmt"
	"time"
)

func main() {

	go count("sheep")

	go count("goast")
	x := 1

	time.Sleep(50 * time.Second)

}

func count(thing string) {

	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)

		time.Sleep(500 * time.Millisecond)
	}

}
