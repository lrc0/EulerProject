//欧拉工程第四题：找出由两个三位数相乘得到的最大的回文数
package main

import (
	"fmt"
	"math"
	// "strconv"
	"time"
)

func main() {
	var result int
	start := time.Now().UnixNano()

	for i := 999; i > 900; i-- {
		for j := 999; j > i; j-- {
			a := i * j
			s := StrNum(a)
			if a == s {
				if s > result {
					result = s
					fmt.Printf("result= [%d], i = [%d], j = [%d]\n", result, i, j)
				}
				break
			}
		}
	}
	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000)
}

//StrNum 数字倒排，例如：1234倒排成4321
func StrNum(s int) int {
	l := len(fmt.Sprint(s))
	var tmp int
	var v int
	for i := l; i > 0; i-- {
		// fmt.Printf("s = [%d]\n", s)
		tmp = s % int(math.Pow10(l-i+1))
		s = s - tmp
		tmp = tmp / int(math.Pow10(l-i))
		v += tmp * int(math.Pow10(i-1))
		// fmt.Printf("tmp = [%d], v = [%d]\n", tmp, v)

	}
	return v
}

// // Str2Num 数字倒排，例如：1234倒排成4321（另一种方法）
// func Str2Num(s int) int {
// 	strs := strconv.Itoa(s)
// 	l := len(strs)
// 	var tmp string
// 	for i := l; i > 0; i-- {
// 		for j, num := range strs {
// 			if j == i-1 {
// 				tmp += string(num)
// 				// fmt.Println(tmp)
// 			}
// 		}
// 	}

// 	intTmp, err := strconv.Atoi(tmp)
// 	if err != nil {
// 		return 0
// 	}
// 	return intTmp
// }
