package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayhello() {
	fmt.Println("Hello, my name is ", p.name, "and my age is ", p.age)
}

func main() {

	p := person{name: "vamshi", age: 22}
	fmt.Println("Hello, my name is ", p.name, "and my age is ", p.age)
}
