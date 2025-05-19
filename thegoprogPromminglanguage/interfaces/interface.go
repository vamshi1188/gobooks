package main

import "fmt"

type printer interface {
	print()
}

type LaserPrinter struct{}

func (lp LaserPrinter) print() {
	fmt.Println("printing with laser printer ")
}

type InkjetPrinter struct{}

func (ip InkjetPrinter) print() {
	fmt.Println("printing with inkjet printer")
}

func main() {

	// var p printer

	// p = LaserPrinter{}

	LaserPrinter{}.print()

	// p = InkjetPrinter{}
	InkjetPrinter{}.print()
}
