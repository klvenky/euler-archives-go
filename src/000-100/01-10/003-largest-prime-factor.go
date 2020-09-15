package main

/*
Largest prime factor
Problem 3
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/
import (
	"fmt"
	"github.com/klvenky/euler-archives-go/src/common"
	"math"
)

func main() {
	maxPrimeFactor := 0
	n := 600851475143
	maxFactor := int(math.Sqrt(float64(n))) + 1
	for factor := 1; factor < maxFactor; factor++ {
		if common.IsPrime(factor) {
			maxPrimeFactor = factor
		}
	}
	fmt.Printf("Biggest prime factor of %d is: %d \n", n, maxPrimeFactor)
}
