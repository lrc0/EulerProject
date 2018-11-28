package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {

	/*
		分析：
			假如是 n X n 的格子，要实现从左上角到右下角的路线，则每一种走法都得横着走n步，再竖着走n步
			所以，如果是20 X 20 的格子，则每一种走法都得横着走20步，再竖着走20步。总共走40步。即在40步
			中取20步横着走，剩下的全竖着走：
				这里适用组合的公式，C(2n, n), 即 40!/(20!*20!)
	*/

	start := time.Now().UnixNano()

	result := new(big.Int)
	fenmu := JieCheng(40)
	fenzi := JieCheng(20)
	fenzi = fenzi.Mul(fenzi, fenzi)
	result = result.Quo(fenmu, fenzi)
	fmt.Println(result)

	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000, "微秒")

}

//JieCheng 计算n的阶乘,涉及到大数运算
func JieCheng(n int64) *big.Int {
	a := new(big.Int)
	y := big.NewInt(1)
	i := int64(1)
	for ; i <= n; i++ {
		x := a.SetInt64(i)
		y = y.Mul(x, y)
	}
	return y
}
