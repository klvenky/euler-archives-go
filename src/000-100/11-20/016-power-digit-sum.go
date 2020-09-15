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

	"github.com/klvenky/euler-archives-go/src/common"
)

func main() {
	// var power uint64 = 15
	var power uint64 = 1000
	two := new(big.Int)
	two.SetUint64(2)
	value := common.CalculatePower(2, power)
	sum := common.DigitSum(value)
	fmt.Printf("Sum of digits of 2^%d is %s\n", power, sum.Text(10))
}
