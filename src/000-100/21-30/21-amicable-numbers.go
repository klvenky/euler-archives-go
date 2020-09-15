/*
Amicable numbers
Problem 21

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
Evaluate the sum of all the amicable numbers under 10000.
*/
package main

import (
	"fmt"
	"math"
)

func sumFactors(n int) int {
	root := int(math.Sqrt(float64(n)))
	sum := 0
	for divisor := 1; divisor <= root; divisor++ {
		if n%divisor == 0 {
			q := n / divisor
			if divisor == root || q == root || q == n {
				sum += divisor
			} else {
				sum += q + divisor
			}
		}
	}
	return sum
}

func main() {
	sum := 0
	max := 10000
	for index := 2; index < max; index++ {
		actualFactorsSum := sumFactors(index)
		sumsFactorSum := sumFactors(actualFactorsSum)
		if index == sumsFactorSum && index != actualFactorsSum {
			sum += index
			fmt.Println("amicable pair received ", index, " & ", sumsFactorSum, ". Sum: ", sum)
		}
	}
	fmt.Printf("Sum of amicable numbers under %d is %d \n", max, sum)
}
