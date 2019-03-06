package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now().UnixNano()

	/*

		由三角形数：1，3，6，10，15 ...
		可知： Sn = S(n-1) + n
			  Sn = (1+n)*n/2
			  S(n-1) = n*(n-1)/2
			  Sn = n*(n-1)/2 + n
			  其中n是自然数1,2,3,4,5 ...
	*/

	s := 1
	for n := 1; ; n++ {
		s = n*(n-1)/2 + n
		if n == 1 || n == 2 {
			continue
		} else {
			result := DecompositionMassFactor(s)
			if result >= 500 && result < 600 {
				fmt.Println("result = ", result)
				fmt.Println("the number is : ", s)
				break
			}
		}
	}

	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000000, "毫秒")
}

//IsPrime 判断质数
func IsPrime(i int) bool {

	if i == 1 {
		return false
	}
	if i == 2 || i == 3 {
		return true
	}

	if i%2 == 0 || i%3 == 0 {
		return false
	}

	if i%6 != 1 && i%6 != 5 {
		return false
	}

	num := int(math.Sqrt(float64(i)))
	for j := 5; j <= num; j += 6 {
		if i%j == 0 || i%(j+2) == 0 {
			return false
		}
	}
	return true
}

//DecompositionMassFactor 分解质因数，并返回x的约数的个数
func DecompositionMassFactor(x int) int {
	var slice []int

	//分解质因数，并将所有质因数放入切片中
outer:
	for {
		for i := 1; i <= x; i++ {

			if i == 1 {
				continue
			} else if x%i == 0 && IsPrime(i) {
				x = x / i
				fmt.Printf("x = %d, i = %d\n", x, i)
				//将所有质因数放入切片中
				slice = append(slice, i)
				if IsPrime(x) {
					//最后剩下的x也是最后一个质因数
					slice = append(slice, x)
					break outer
				}
			}
		}
	}
	fmt.Println("slice: ", slice)
	//统计相同的质因数的个数
	m := make(map[int]int)
	for _, num := range slice {
		//质因数的个数，key是质因数，value是质因数的个数
		m[num]++
	}

	count := 1
	fmt.Println("map :", m)
	/*
		计算所有的约数，根据约数个数定理：
			首先同上，n可以分解质因数：n=p1^a1×p2^a2×p3^a3*…*pk^ak,
		由约数定义可知p1^a1的约数有:p1^0, p1^1, p1^2......p1^a1 ，共（a1+1）个;同理p2^a2的约数有（a2+1）个......pk^ak的约数有（ak+1）个。
		故根据乘法原理：n的约数的个数就是(a1+1)(a2+1)(a3+1)…(ak+1)。

	*/
	for _, v := range m {
		count *= v + 1
	}
	//count是约数的个数
	fmt.Println("count = ", count)
	return count
}
