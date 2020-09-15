package common

import (
	"fmt"
	"math"
	"math/big"
)

// IsPrime Checks if a function is prime or not
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

// IsFactor is a function that returns true/fa;se based on the two inputs
func IsFactor(n int, factor int) bool {
	return (n % factor) == 0
}

// DigitSum adds the individual digits of a long number
func DigitSum(value big.Int) big.Int {
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
