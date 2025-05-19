package main

import "fmt"

type payable interface {
	pay(amount float64)
}

type PayPal struct{}
type CreditCard struct{}

func (pp PayPal) pay(amount float64) {
	fmt.Printf("paid $%.2f with paypal \n", amount)
}

func (cc CreditCard) pay(amount float64) {
	fmt.Printf("paid $%.2f with credit card \n", amount)
}

func makepayment(p payable, amt float64) {

	p.pay(amt)
}

func main() {

	cc := CreditCard{}
	pp := PayPal{}

	makepayment(cc, 45.2)
	makepayment(pp, 22.2)
}
