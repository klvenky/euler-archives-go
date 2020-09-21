package main

/*
Largest palindrome product
Problem 4
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/

import (
	"fmt"

	"github.com/klvenky/euler-archives-go/src/common"
)

func main() {
	// Example problem values
	// max := 99
	// min := 10
	// Expected output: 9009

	max := 999
	min := 100
	result := 0
	for i := max; i >= min; i-- {
		for j := max; j >= min; j-- {
			product := i * j
			isPal := common.IsPalindrome(product)
			if product > result && isPal {
				result = product
			}
		}
	}
	fmt.Println("Max palindrome product is ", result)
}
