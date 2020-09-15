/*
Factorial digit sum
Problem 20
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/
package main

import (
	"fmt"
	"github.com/klvenky/euler-archives-go/src/common"
	"math/big"
)

func factorial(number int64) big.Int {
	one := new(big.Int)
	one.SetInt64(1)
	input := new(big.Int)
	input.SetInt64(number)
	bigInt := new(big.Int)
	bigInt.SetInt64(number)
	if number <= 1 {
		fmt.Print("returning 1\n")
		bigInt.SetInt64(1)
	} else {
		nextVal := factorial(number - 1)
		bigInt = bigInt.Mul(bigInt, &nextVal)
	}
	return *bigInt
}

func main() {
	var n int64 = 100
	fact := factorial(n)
	fmt.Println(&fact)
	fmt.Printf("Sum of digits in factorial(%d) is %d\n", n, common.DigitSum(fact))
}
