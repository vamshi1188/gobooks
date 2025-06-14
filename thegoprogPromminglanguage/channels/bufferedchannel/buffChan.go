package main

import "fmt"

func main() {

	ch := make(chan string, 3)

	ch <- "a"
	ch <- "b"
	ch <- "c"
	fmt.Println(<-ch)

	ch <- "d"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
