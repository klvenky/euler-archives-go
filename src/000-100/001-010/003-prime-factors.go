package main

/*
Largest prime factor
Problem 3
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/
import (
	"fmt"
	"math"
)

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

func isFactor(n int, factor int) bool {
	return (n % factor) == 0
}

func main() {
	maxPrimeFactor := 0
	n := 600851475143
	maxFactor := int(math.Sqrt(float64(n))) + 1
	for factor := 1; factor < maxFactor; factor++ {
		if isFactor(n, factor) && isPrime(factor) {
			maxPrimeFactor = factor
		}
	}
	fmt.Printf("Biggest prime factor of %d is: %d \n", n, maxPrimeFactor)
}
