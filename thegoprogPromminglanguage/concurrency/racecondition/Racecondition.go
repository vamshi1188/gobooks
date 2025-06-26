package main

import (
	"fmt"
	"time"
)

var balances = make(chan int)
var deposits = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
func main() {

	go func() {
		Deposit(200)
		fmt.Println("alice sees balance = ", Balance())
	}()

	go func() {
		Deposit(100)
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("final balance :", Balance())
}
