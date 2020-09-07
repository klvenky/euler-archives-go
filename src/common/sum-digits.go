import (
	"big"
)

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
