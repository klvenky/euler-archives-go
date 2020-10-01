/*
Problem 30 - Digit fifth powers
-----------
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

    1634 = 1^4 + 6^4 + 3^4 + 4^4
    8208 = 8^4 + 2^4 + 0^4 + 8^4
    9474 = 9^4 + 4^4 + 7^4 + 4^4

As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

package main

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/klvenky/euler-archives-go/src/common"
)

func main() {
	const max = 999999
	const min = 2
	const power = 5
	matches := make([]int, 0)
	sum := 0
	// bigNum := common.GetDigits(common.CalculatePower(10, 3))
	for i := min; i <= max; i++ {
		bigNum := new(big.Int)
		bigNum.SetInt64(int64(i))
		digits := common.GetDigits(*bigNum)
		digitPowerSum := new(big.Int)
		digitPowerSum.SetInt64(0)
		for _, item := range digits {
			powerVal := common.CalculatePower(uint64(item), uint64(power))
			digitPowerSum.Add(digitPowerSum, &powerVal)
		}
		if digitPowerSum.Text(10) == bigNum.Text(10) {
			intNum, _ := strconv.Atoi(digitPowerSum.Text(10))
			matches = append(matches, intNum)
		}

	}
	// fmt.Println(matches)
	for _, match := range matches {
		sum += match
	}
	fmt.Printf("Sum of all numbers that can be written as %dth powers of their digits is %d\n", power, sum)
}
