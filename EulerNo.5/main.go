package main

import (
	"fmt"
	"math"
	"time"
)

/*
	分析:

	20以内的质数：2, 3, 5, 7, 11, 13, 17, 19
	剩下: 4, 6, 8, 10, 12, 14, 16, 18, 20

		剩下的合数:
		4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20
		拆分成质数相乘:
			4  == 2 * 2
			8  == 2 * 2 * 2
			9  == 3 * 3
			10 == 2 * 5
			12 == 2 * 2 * 3
			14 == 2 * 7
			15 == 3 * 5
			16 == 2 * 2 * 2 * 2
			18 == 2 * 3 * 3
			20 == 2 * 2 * 5

	最小公倍数 = 2 * 2 * 2 * 3 * 2 * 3 * 5 * 7 = 5040
	因为在上一步已经包括进去2, 3, 5, 7这4个质数
	所以1-20内最小公倍数 = 2 * 2 * 2 * 3 * (2 * 3 * 5 * 7 * 11 * 13 * 17 * 19) = 232792560


*/

//代码实现
func main() {
	start := time.Now().UnixNano()

	target := 20

	//算出target以内所有质数的积product, 收集target以内除去质数剩下的合数集合s
	product, s := FindPrimeNumber(target)
	//算出集合s内的所有合数的最小公倍数
	total := FindCommonMultiple(s)

	//去除最小公倍数内已使用过一次的target以内包含的质数
	num := RemovePrimeCommonDivisor(target, total)

	//最终结果
	result := product * num
	fmt.Println(result)

	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000000)

}

//FindPrimeNumber 找出小于等于a的所有素数,并将他们相乘
func FindPrimeNumber(target int) (int, []int) {
	var slice []int
	num := 1

	for i := 2; i <= target; i++ {
		if IsPrime(i) {
			fmt.Println("target 以内的质数：", i)
			num *= i //将所有质数相乘
		} else {
			slice = append(slice, i) //收集剩下的合数
		}
	}
	return num, slice
}

//FindCommonMultiple 求这些合数的最小公倍数
func FindCommonMultiple(slice []int) int {
	total := 1
	for _, num := range slice {
		total = MinCommonMultiple(total, num)
	}
	return total
}

//RemovePrimeCommonDivisor 去除一次合数的最小公倍数里面已使用过的质数
func RemovePrimeCommonDivisor(target, num int) int {
	for i := 2; i <= target; i++ {
		if IsPrime(i) {
			if num%i == 0 {
				num = num / i
			}
		}
	}
	return num
}

//MaxCommonDivisor 最大公约数
func MaxCommonDivisor(x, y int) int {
	var n int
	if x > y {
		n = x
	} else {
		n = y
	}
	for i := n; i > 0; i-- {
		if x%i == 0 && y%i == 0 {
			return i
		}
	}
	return 1
}

//MinCommonMultiple 最小公倍数, 公式解法
func MinCommonMultiple(x, y int) int {
	total := x * y
	divisor := MaxCommonDivisor(x, y)

	// 由于两个数的乘积等于这两个数的最大公约数与最小公倍数的积:即（a,b）*[a,b]=a×b,所以最小公倍数[a,b]=a*b/(a,b)
	multiple := total / divisor
	return multiple
}

/*
	在网上看到一种高效的方法, 连接:http://blog.csdn.net/huang_miao_xin/article/details/51331710

	一个数若可以进行因数分解，那么分解时得到的两个数一定是一个小于等于sqrt(n)，一个大于等于sqrt(n)
	所以遍历的时候只需要遍历到sqrt(n)即可

	首先看一个关于质数分布的规律：大于等于5的质数一定和6的倍数相邻。例如5和7，11和13,17和19等等；

	证明：令x≥1，将大于等于5的自然数表示如下：
	······ 6x-1，6x，6x+1，6x+2，6x+3，6x+4，6x+5，6(x+1），6(x+1)+1 ······
	可以看到，不在6的倍数两侧，即6x两侧的数为6x+2，6x+3，6x+4，由于2(3x+1)，3(2x+1)，2(3x+2)，
	所以它们一定不是素数，再除去6x本身，显然，素数要出现只可能出现在6x的相邻两侧。这里有个题外话，关于孪生素数，
	有兴趣的道友可以再另行了解一下，由于与我们主题无关，暂且跳过。这里要注意的一点是，在6的倍数相邻两侧并不是一定就是质数。
	根据以上规律，判断质数可以6个为单元快进，即将方法（2）循环中i++步长加大为6，加快判断速度，代码如下：
*/

//IsPrime 判断是否是质数
func IsPrime(i int) bool {
	if i == 2 || i == 3 {
		return true
	}

	//判断i能否被2或3整除
	if i%2 == 0 || i%3 == 0 {
		return false
	}

	//不在6两侧的数,不是质数(并不一定,不过至少100以内适用)
	if i%6 != 1 && i%6 != 5 {
		if !IsPrime(i) {
			fmt.Println("在6的倍数的两侧不是素数")
			return false
		}
		return false
	}

	//只需要遍历到sqrt(i)即可
	x := float64(i)
	n := int(math.Sqrt(x))
	for j := 5; j <= n; j += 6 {
		if i%j == 0 || i%(j+2) == 0 {
			return false
		}
	}

	return true
}
