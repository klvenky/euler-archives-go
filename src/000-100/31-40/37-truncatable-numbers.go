package main

/*
Truncatable primes
Problem 37
The number 3797 has an interesting property. Being prime itself, it is possible to continuously
remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7.

Similarly we can work from right to left: 3797, 379, 37, and 3.
Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/
import (
	"fmt"
	"math/big"

	"github.com/klvenky/euler-archives-go/src/common"
)

func isOneWayTruncatablePrime(num int, first bool, removeFromRight bool) bool {
	if common.IsPrime(num) {
		bigNum := new(big.Int)
		bigNum.SetInt64(int64(num))
		digits := common.GetDigits(*bigNum)
		singleDigit := len(digits) == 1
		if singleDigit {
			if first {
				return false
			} else {
				isSingleDigitPrime := common.IsPrime(num)
				return isSingleDigitPrime
			}
		} else {
			genNum := 0
			for index, digit := range digits {
				if removeFromRight && index != 0 {
					genNum *= 10
					genNum += digit
				} else if !removeFromRight && index != len(digits)-1 {
					genNum *= 10
					genNum += digit
				}

			}

			oneWayP := isOneWayTruncatablePrime(genNum, false, removeFromRight)
			return oneWayP

		}
	}
	return false
}
func isTruncatablePrime(num int) bool {
	bigNum := new(big.Int)
	bigNum.SetInt64(int64(num))
	rightPrime := isOneWayTruncatablePrime(num, true, true)
	if rightPrime {
		leftPrime := isOneWayTruncatablePrime(num, true, false)
		return leftPrime
	} else {
		return rightPrime
	}
}

func main() {
	sum := 0
	max := 11
	for counter, index := 0, 2; counter < max; index += 1 {
		if isTruncatablePrime(index) {
			fmt.Println(index, " is a truncatable prime ")
			sum += index
			counter += 1
		}

	}
	fmt.Println("sum of 11 truncatable primes is ", sum)
	//fmt.Println("Truncatable Prime 3797 : ", isTruncatablePrime(3797))
}
