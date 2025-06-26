package main

import (
	"fmt"
	"sync"
	"time"
)

var x, y int
var mu sync.Mutex

func main() {
	go func() {
		mu.Lock()
		x = 1
		fmt.Println("y:", y)
		mu.Unlock()
	}()

	go func() {
		mu.Lock()
		y = 1
		fmt.Println("x:", x)
		mu.Unlock()
	}()
	time.Sleep(1 * time.Second)
}
