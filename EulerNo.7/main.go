package main

import (
	"fmt"
	"math"
	"time"
)

/*

	分析:

	在网上看到一种高效的方法, 连接:http://blog.csdn.net/huang_miao_xin/article/details/51331710

	一个数若可以进行因数分解，那么分解时得到的两个数一定是一个小于等于sqrt(n)，一个大于等于sqrt(n)
	所以遍历的时候只需要遍历到sqrt(n)即可
	首先看一个关于质数分布的规律：大于等于5的质数一定和6的倍数相邻。例如5和7，11和13,17和19等等；
	证明：令x≥1，将大于等于5的自然数表示如下：
	······ 6x-1，6x，6x+1，6x+2，6x+3，6x+4，6x+5，6(x+1），6(x+1)+1 ······
	可以看到，不在6的倍数两侧，即6x两侧的数为6x+2，6x+3，6x+4，由于2(3x+1)，3(2x+1)，2(3x+2)，
	所以它们一定不是素数，再除去6x本身，显然，素数要出现只可能出现在6x的相邻两侧。这里有个题外话，关于孪生素数，
	有兴趣的道友可以再另行了解一下，由于与我们主题无关，暂且跳过。这里要注意的一点是，在6的倍数相邻两侧并不是一定就是质数。
	根据以上规律，判断质数可以6个为单元快进，即将方法（2）循环中i++步长加大为6，加快判断速度

	题干要求找出第10001个质数, 所以笨办法是从100000开始试, 100000为单位往上加

*/

func main() {
	start := time.Now().UnixNano()
	var slice []int
	for i := 2; i < 200000; i++ {
		if IsPrime(i) {
			slice = append(slice, i)
		}
	}

	for i, num := range slice {
		if i == 10001 {
			fmt.Println("第10001个质数:", num)
			break
		}
	}
	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000000)
}

//IsPrime 判断是否是素数
func IsPrime(i int) bool {
	if i == 2 || i == 3 {
		return true
	}
	if i%2 == 0 || i%3 == 0 {
		return false
	}

	//不在6的倍数两侧的数,不是质数(并不一定,不过至少100以内适用)
	if i%6 != 1 && i%6 != 5 {
		if !IsPrime(i) {
			fmt.Println("在6的倍数的两侧，但不是素数")
			return false
		}
		return false
	}

	var n = float64(i)
	x := int(math.Sqrt(n))
	for j := 5; j <= x; j += 6 {
		if i%j == 0 || i%(j+2) == 0 {
			return false
		}
	}
	return true
}
