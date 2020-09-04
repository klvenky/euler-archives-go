/*
Power digit sum
Problem 16
2 Power 15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2 power 1000?
*/

package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func main() {
	power := 15
	// power := 1000
	value := new(big.Int)
	powerCaclulated := math.Pow(float64(2), float64(power))
	
	// value.SetString(strconv.FormatInt(powerCaclulated, 64), 10)
	fmt.Printf("%dth power of 2 : %d\n", power, value)
	sum := new(big.Int)
	sum.SetUint64(0)
	zero := new(big.Int)
	zero.SetUint64(0)
	ten := new(big.Int)
	ten.SetUint64(10)
	for i := value; i.Cmp(zero) > 0; {
		rem := new(big.Int)
		q := new(big.Int)
		q, rem = i.DivMod(i, ten, rem)
		fmt.Println(q, i)
		sum.Add(sum, rem)
		fmt.Printf("quotient: %d | Remainder:%d | Dividend: %d | Sum: %d\n", q, rem, i, sum)
		i = q
	}
	fmt.Printf("sum of digits in %d power of 2 is %d\n", power, sum)
}
