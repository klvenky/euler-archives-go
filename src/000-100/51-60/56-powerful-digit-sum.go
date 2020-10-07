package main

import (
	"fmt"

	"github.com/klvenky/euler-archives-go/src/common"
)

/*
Powerful digit sum
Problem 56

A googol (10^100) is a massive number: one followed by one-hundred zeros;
100^100 is almost unimaginably large: one followed by two-hundred zeros.
Despite their size, the sum of the digits in each number is only 1.

Considering natural numbers of the form, ab, where a, b < 100, what is the maximum digital sum?
*/

func main() {
	var max uint64 = 100
	maxSum := 0
	var a uint64
	for a = 1; a < max; a++ {
		var b uint64
		for b = 1; b < max; b++ {
			tmpSum := 0
			for _, digit := range common.GetDigits(common.CalculatePower(a, b)) {
				tmpSum += digit
			}
			if tmpSum > maxSum {
				maxSum = tmpSum
			}
		}
	}
	fmt.Println("max powerful digit sum is ", maxSum)
}
