/*
Circular primes
Problem 35

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
How many circular primes are there below one million?
*/
package main

import (
	"fmt"
	"math/big"

	"github.com/klvenky/euler-archives-go/src/common"
)

func getCircularNumbers(digits []int) []int {
	if len(digits) == 1 {
		return digits
	} else {
		clone := make([]int, len(digits))
		circularNums := make([]int, len(digits))
		copy(clone, digits)
		for i := 0; i < len(digits); i++ {
			d, copy := clone[0], clone[1:]
			copy = append(copy, d)
			num := 0
			for i := 0; i < len(copy); i++ {
				num = num*10 + copy[i]
			}
			clone = copy
			circularNums[i] = num
		}
		return circularNums
	}
}

func main() {
	// max := 100
	const max int = 1000000
	count := 0
	cps := make([]int, max)
	for i := 2; i < max; i++ {
		if common.IsPrime(i) {
			curr := new(big.Int)
			curr.SetUint64(uint64(i))
			digits := common.GetDigits(*curr)
			if len(digits) == 1 {
				count++
				cps[count-1] = i
			} else {
				circularNums := getCircularNumbers(digits)
				cNumsCount := len(circularNums)
				primes := 0

				for _, num := range circularNums {
					prim := common.IsPrime(num)
					if prim {
						primes++
					}
				}
				if primes == cNumsCount {
					count++
					cps[count-1] = i
				}
			}
		}
	}
	fmt.Printf("count is %d\n", count)
	// fmt.Println(cps[0:count])
}
