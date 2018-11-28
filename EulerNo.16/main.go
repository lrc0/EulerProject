package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func main() {

	start := time.Now().UnixNano()

	x := big.NewInt(2)
	y := big.NewInt(1000)
	z := new(big.Int)

	z = z.Exp(x, y, nil)
	zStr := z.String()

	var result int

	for _, numRune := range zStr {
		numStr := string(numRune)
		numInt, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		result += numInt
	}
	fmt.Println(result)

	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000, "微秒")
}
