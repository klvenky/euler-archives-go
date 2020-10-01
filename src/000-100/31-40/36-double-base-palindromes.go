/*
Double-base palindromes
Problem 36
The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.
Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
(Please note that the palindromic number, in either base, may not include leading zeros.)
*/

package main

import (
	"fmt"
	"strconv"

	"github.com/klvenky/euler-archives-go/src/common"
)

func isDoublePalindrome(number int) bool {
	if common.IsPalindrome(number) {
		binaryString := strconv.FormatInt(int64(number), 2)
		// fmt.Println(number, " in binary format is ", binaryString)
		isDP := common.IsStrPalindrome(binaryString)
		if isDP {
			fmt.Println(number, "is plaindrome ")
		}
		return isDP
	}
	return false
}

func main() {
	sum := 0
	const max int = 1000000
	for i := 1; i < max; i += 1 {
		if isDoublePalindrome(i) {
			sum += i
		}
	}
	fmt.Println("sum of all double palindromes under ", max, " is ", sum)
}
