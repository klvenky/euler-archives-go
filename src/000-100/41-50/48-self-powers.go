package main

import (
	"fmt"
	"math/big"

	"github.com/klvenky/euler-archives-go/src/common"
)

/*
Self powers
Problem 48

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.
Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
*/

func main() {
	var max uint64 = 1000
	// result := ""
	sum := new(big.Int)
	sum.SetUint64(0)
	var num uint64
	for num = 1; num <= max; num++ {
		bigPow := common.CalculatePower(num, num)
		sum.Add(sum, &bigPow)
	}
	fmt.Printf("last 10 digits of self powers upto %d is %s", max, sum.Text(10))
}
