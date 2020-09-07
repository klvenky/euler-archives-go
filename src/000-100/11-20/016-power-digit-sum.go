/*
Power digit sum
Problem 16
2 Power 15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2 power 1000?
*/

package main

import (
	"fmt"
	"math/big"
)

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
		// fmt.Println(q, i)
		sum.Add(sum, rem)
		// fmt.Printf("quotient: %d | Remainder:%d | Dividend: %d | Sum: %d\n", q, rem, i, sum)
		i = *q
	}
	fmt.Println("digit sum is ", sum)
	return *sum
}

func calculatePower(number uint64, power uint64) big.Int {
	bigNum := new(big.Int)
	bigNum.SetUint64(number)
	powerBig := new(big.Int)
	powerBig.SetUint64(power)
	result := new(big.Int)
	if power < 1 {
		result.SetUint64(1)
	} else if power == 1 {
		result.SetUint64(number)
		fmt.Println(result)
	} else {
		result.SetUint64(1)
		// fmt.Println("Initial value of result is : ", result, " power is :", power)
		for index := 1; uint64(index) <= power; index++ {
			result = result.Mul(result, bigNum)
			// fmt.Println(index, " is ", index, ", result is ", result)
		}
	}
	fmt.Println("result is : ", result.Text(10))
	return *result
}

func main() {
	// var power uint64 = 15
	var power uint64 = 1000
	two := new(big.Int)
	two.SetUint64(2)
	value := calculatePower(2, power)
	sum := digitSum(value)
	fmt.Printf("Sum of digits of 2^%d is %s\n", power, sum.Text(10))
}
