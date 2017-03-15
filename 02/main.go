package main

import (
	"fmt"
)

// Ex02 contains the logic needed to produce the required output for this exercise
func Ex02(limit int) (result int) {
	result = 0
	prev := 1
	curr := 1

	for i := 1; i < limit; i = prev + curr {
		prev = curr
		curr = i
		if i%2 == 0 {
			result += i
		}
	}

	return result
}

func main() {
	fmt.Println("Ex02", Ex02(4000000))
}
