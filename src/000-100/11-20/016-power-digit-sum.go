/*
Power digit sum
Problem 16
2 Power 15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2 power 1000?
*/

package main

import (
	"fmt"

	"github.com/klvenky/euler-archives-go/src/common"
)

func main() {
	// var power uint64 = 15
	var index uint64 = 2
	var power uint64 = 1000
	calculatedPowerVal := common.CalculatePower(index, power)
	sum := common.DigitSum(calculatedPowerVal)
	fmt.Printf("Sum of digits of 2^%d is %d\n", power, sum)
}
