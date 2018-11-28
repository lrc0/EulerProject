package main

import (
	"fmt"
	"strings"
	"time"
)

/*
	分析：
		1 - 9: one two three four five six seven eight nine
		10-19: ten eleven twelve thirteen fourteen fifteen sixteen seventeen eighteen nineteen


		百,千,百万,10亿: Hundred,Thousand,Million,Billion

			Billion = 1000 * Million
			Million = 1000 * Thousand

		在1-99里面可以分成几段: 1-9, 10-19, 20,30,40,50,60,70,80,90 和剩下的

		可以看出，上面列出来的已经包含了所有的能变化的数字。其他数字全是由这些英文组合起来的

		由此思路:
			1.将所有数字翻译成英文形式
			2.遍历1-1000,并将这些数字翻译成英文
			3.统计这些英文的字母数量

	PS: 其实这题很无聊, 稍微统计一下就能计算出来。这里完全是为了学习使用golang, 才绕一个大弯写一个翻译程序来处理。

*/

func main() {

	start := time.Now().UnixNano()

	count := 0
	for i := 1; i <= 1000; i++ {
		x := NumberToEnglish(i)
		count += CountTheWord(x)
	}
	fmt.Println("result = ", count)

	// for i := 100000; i <= 999999; i++ {
	// 	x := NumberToEnglish(i)
	// 	fmt.Printf("[%d] = %s\n", i, x)
	// }

	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000000, "毫秒")
}

func CountTheWord(s string) int {
	var count int
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, " ", "", -1)
	count = len(s)
	return count
}

//NumberToEnglish 数字翻译成英文
func NumberToEnglish(x int) string {
	var y string
	var num int

	m := make(map[int]string)
	//1-9
	num1 := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	// 10 - 19
	num2 := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	//20,30,40,50,60,70,80,90
	num3 := []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	//1-19
	for i, w := range num1 {
		m[i+1] = w
	}
	for i, w := range num2 {
		m[i+10] = w
	}

	if x > 0 && x <= 19 {
		y = m[x]
	}

	//20-99, 闭包来处理, 以便后面复用
	fn := func(x int) string {
		var a string
		n := x / 10
		n = n - 2
		if x%10 == 0 {
			a = num3[n]
		} else {
			a = num3[n] + "-" + m[x%10]
		}
		return a
	}

	//1-99
	fnx := func(x int) string {
		var res string

		if x > 0 && x <= 19 {
			res = m[x]
		} else if x >= 20 && x <= 99 {
			res = fn(x)
		}
		return res
	}

	//100-999, 闭包来处理, 以便后面复用
	fn1 := func(x int) string {
		var b string
		num = x / 100
		if x%100 == 0 {
			b = m[num] + " hundred"
		} else {
			a := x % 100
			b = m[num] + " hundred and " + fnx(a)
		}
		return b
	}

	//1-999
	fnz := func(x int) string {
		var res string
		if x > 0 && x <= 19 {
			res = m[x]
		} else if x >= 20 && x <= 99 {
			res = fn(x)
		} else if x >= 100 && x <= 999 {
			res = fn1(x)
		}
		return res
	}

	//1000-999999
	if x >= 1000 && x <= 999999 {
		s := x / 1000
		w := x % 1000
		if w == 0 {
			y = fnz(s) + " thousand"
		} else if w >= 1 && w <= 999 {
			y = fnz(s) + " thousand and " + fnz(w)
		}
	}
	//todo Billion

	if x >= 1 && x <= 999 {
		y = fnz(x)
	}

	return y
}
