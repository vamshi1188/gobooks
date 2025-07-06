// Package word provides utilities for word games.
package main

import "fmt"

// IsPalindrome reports whether s is a palindrome.
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	anser := IsPalindrome("sasd")

	fmt.Println(anser)
}
