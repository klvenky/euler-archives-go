/*
Distinct primes factors
Problem 47

The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?
*/

package main

import (
	"math"

	"github.com/klvenky/euler-archives-go/src/common"
)

func main() {
	var primes []int
	const maxFactors int = 2
	for numOne := 3; ; numOne++ {
		sqrt := int(math.Sqrt(float64(numOne)))
		for factor := 2; factor < sqrt; factor++ {
			isFactor := common.IsFactor(numOne, factor)
			if isFactor && common.IsPrime(factor) {
				primes = append(primes, factor)
			}
		}

	}

}
