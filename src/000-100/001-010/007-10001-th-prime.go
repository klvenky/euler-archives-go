package main

/* 10001st prime
Problem 7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
*/
import (
	"fmt"
	"math"
	"time"
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
func main() {
	// primeIndex := 6
	primeIndex := 10001
	result := 0
	start := time.Now()
	for counter, i := 0, 1; counter < primeIndex; i++ {

		if isPrime(i) {
			counter++
			result = i
		}
	}
	timeTaken := time.Since(start)
	fmt.Printf("%dth Prime number is %d\n", primeIndex, result)
	fmt.Println("took ", timeTaken)
}
