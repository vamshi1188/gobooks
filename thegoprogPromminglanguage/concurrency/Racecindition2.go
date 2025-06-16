package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(" Race - condition")

	mut := sync.Mutex{}

	var score = []int{0}
	go func() {
		fmt.Println("one R")
		mut.Lock()
		defer mut.Unlock()
		score = append(score, 1)
		fmt.Println(score)

	}()
	go func() {
		fmt.Println("two R")
		mut.Lock()
		defer mut.Unlock()
		score = append(score, 2)
		fmt.Println(score)

	}()
	go func() {
		fmt.Println("three R")
		mut.Lock()
		defer mut.Unlock()
		score = append(score, 3)
		fmt.Println(score)

	}()

	//time.Sleep(1 * time.Second)
}
