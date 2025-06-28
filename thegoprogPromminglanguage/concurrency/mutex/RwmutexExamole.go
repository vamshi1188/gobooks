package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu      sync.RWMutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Depositing:", amount)
	balance += amount
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}

func main() {
	go Deposit(100)

	for i := 1; i <= 5; i++ {
		go func(id int) {
			fmt.Printf("Reader %d: Balance is %d\n", id, Balance())
		}(i)
	}

	time.Sleep(2 * time.Second)
}
