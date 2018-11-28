package main

import (
	"fmt"
	"math"
)

/*
	分析:

	先算出1的平方连加到100的平方得出total
	再算出1加到100的和, 再将其平方, 得到total1
	结果result = total1 - total


*/

func main() {
	var total float64
	var total1 float64
	a := 0
	for i := 1; i <= 100; i++ {
		total += math.Pow(float64(i), 2)
		a += i
	}
	total1 += math.Pow(float64(a), 2)
	result := total1 - total
	fmt.Println(result)
}
