package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().UnixNano()

	var result int
	var i int
	for n := 2; n < 1000000; n++ {
		i = n
		count := Count(i)
		if result < count {
			result = count
			fmt.Printf("len = %d, the number is : %d\n", result, i)
		}
	}

	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000000, "毫秒")
}

//Count 返回计算步骤长度
func Count(i int) int {
	count := 0
	for {
		if i%2 != 0 {
			i = i*3 + 1
			// fmt.Println("i = ", i)
			count++

		} else if i%2 == 0 {
			i = i / 2
			// fmt.Println("i = ", i)
			count++
			if i == 1 {
				// fmt.Println("count = ", count)
				break
			}
		}
	}
	return count
}
