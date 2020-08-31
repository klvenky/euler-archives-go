package main

/*
Largest palindrome product
Problem 4
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for index := 0; index < length-1; index++ {
		if int(str[index]) != int(str[length-index-1]) {
			return false
		}
	}
	return true
}

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
			isPal := isPalindrome(product)
			if product > result && isPal {
				result = product
			}
		}
	}
	fmt.Println("Max palindrome product is ", result)
}
