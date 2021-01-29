package main

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/klvenky/euler-archives-go/src/common"
)

/*
Permuted multiples
Problem 52
It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.
Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.
*/
func sortDigits(digits []int) {
	sort.Ints(digits)
}
func main() {
	result := 0
	for num := 1; ; num++ {
		count := 0
		for multiplier := 2; multiplier <= 6; multiplier++ {
			if checkSameDigitsForFactor(uint64(num), uint64(multiplier)) {
				count++
			} else {
				break
			}
		}
		if count == 6 {
			result = num
			break
		} else {
			if num%100000 == 0 {
				fmt.Println("Done till ", num)
			}
		}
	}
	fmt.Println("the smallest permuted multiple is : ", result)
}

func checkSameDigitsForFactor(num uint64, multiplyBy uint64) bool {
	bigNum := new(big.Int)
	bigNum.SetUint64(num)
	numDigits := common.GetDigits(*bigNum)
	sort.Ints(numDigits)
	product := num * multiplyBy
	productBigNum := new(big.Int)
	productBigNum.SetUint64(product)
	prodDigits := common.GetDigits(*productBigNum)
	return common.AreIntArrayEqual(numDigits, prodDigits)
}
