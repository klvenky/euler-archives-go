import (
	"big"
	"fmt"
)

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
