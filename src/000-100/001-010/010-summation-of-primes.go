package main

import (
	"fmt"
	"math"
)

/*
Summation of primes
Problem 10
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
*/

func isPrime(n int) bool {
	out := false
	if n == 2 {
		out = true
	} else if (n % 2) == 0 {
		out = false
	} else if n < 2 {
		out = false
	} else {
		max := int(math.Sqrt(float64(n)))
		factors := 1
		for factor := 2; factor < max+1; factor = factor + 1 {
			if n%factor == 0 {
				factors++
			}
			if factors > 1 {
				break
			}
		}
		if factors < 2 {
			out = true
		}
	}
	return out
}

func main() {
	// max := 10
	max := 2000000
	sum := 0
	for i := 2; i < max; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Printf("Sum of primes below %d is %d\n", max, sum)
}
