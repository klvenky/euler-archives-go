package main

import (
	"fmt"
	"math"
)

func isPrime(n int32) bool {
	out := false
	if n == 2 {
		out = true
	} else if (n % 2) == 0 {
		fmt.Println("n is 2")
		out = false
	} else if n < 2 {
		out = false
		fmt.Println("n is lessthan 2")
	} else {
		fmt.Println("in else")
		max := math.Sqrt(float64(n))
		fmt.Printf("printed %f", max)
	}
	fmt.Println(out)
	return out
}
