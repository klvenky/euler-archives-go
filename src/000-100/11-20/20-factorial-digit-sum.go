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

func digitSum(value big.Int) big.Int {
	sum := new(big.Int)
	sum.SetUint64(0)
	zero := new(big.Int)
	zero.SetUint64(0)
	ten := new(big.Int)
	ten.SetUint64(10)

	for i := value; i.Cmp(zero) > 0; {
		rem := new(big.Int)
		q := new(big.Int)
		q, rem = i.DivMod(&i, ten, rem)
		fmt.Println(q, i)
		sum.Add(sum, rem)
		// fmt.Printf("quotient: %d | Remainder:%d | Dividend: %d | Sum: %d\n", q, rem, i, sum)
		i = *q
	}
	return *sum
}

func main() {
	var n int64 = 100
	fact := factorial(n)
	fmt.Println(&fact)
	fmt.Printf("Sum of digits in factorial(%d) is %d\n", n, digitSum(fact))
}
