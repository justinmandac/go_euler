package main

import (
	"fmt"
)

// Create a program that will get the sum of all multiples of 3 and 5 below 1000

// Ex01 contains logic for Project Euler Exercise # 1
func Ex01(input int, limit int) (result int) {
	result = 0

	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}

	return result
}

func main() {
	fmt.Println("Ex01 :", Ex01(1, 10000))
}
