package common

import (
	"math"
	"math/big"
	"strconv"
)

// IsPrime Checks if the input number is prime or not
func IsPrime(n int) bool {
	out := false
	if n == 2 {
		out = true
	} else if (n % 2) == 0 {
		out = false
	} else if n < 2 {
		out = false
	} else {
		max := int(math.Sqrt(float64(n)))
		factors := 1
		for factor := 2; factor < max+1; factor = factor + 1 {
			if n%factor == 0 {
				factors++
			}
			if factors > 1 {
				break
			}
		}
		if factors < 2 {
			out = true
		}
	}
	return out
}

// IsFactor returns true/false based on the two inputs
func IsFactor(n int, factor int) bool {
	return (n % factor) == 0
}

// DigitSum adds the individual digits of a long number
func DigitSum(value big.Int) int {
	digits := GetDigits(value)
	sum := 0
	for _, digit := range digits {
		sum += digit
	}
	return sum
}

// CalculatePower takes the root and the power to calculate
func CalculatePower(number uint64, power uint64) big.Int {
	bigNum := new(big.Int)
	bigNum.SetUint64(number)
	powerBig := new(big.Int)
	powerBig.SetUint64(power)
	result := new(big.Int)
	if power < 1 {
		result.SetUint64(1)
	} else if power == 1 {
		result.SetUint64(number)
	} else {
		result.SetUint64(1)
		for index := 1; uint64(index) <= power; index++ {
			result = result.Mul(result, bigNum)
		}
	}
	return *result
}

func FindDigitCount(number big.Int) int {
	digits := GetDigits(number)
	return len(digits)
}

// GetDigits gives an array of integers that are part of a supplied bigInt number
func GetDigits(number big.Int) []int {
	var nums []int
	zero := new(big.Int)
	zero.SetUint64(0)
	ten := new(big.Int)
	ten.SetUint64(10)
	one := new(big.Int)
	one.SetUint64(1)
	tmp := new(big.Int)
	tmp.SetUint64(0)
	tmp = tmp.Add(tmp, &number)
	for i := tmp; i.Cmp(zero) > 0; {
		rem := new(big.Int)
		_, rem = i.DivMod(i, ten, rem)
		v, _ := strconv.Atoi(rem.Text(10))
		nums = append(nums, v)
	}
	digits := make([]int, len(nums))
	numsLen := len(nums)
	for i := 0; i < numsLen; i += 1 {
		digits[i] = nums[numsLen-i-1]
	}
	return digits
}

func IsPalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for index := 0; index < length-1; index++ {
		if int(str[index]) != int(str[length-index-1]) {
			return false
		}
	}
	return true
}
