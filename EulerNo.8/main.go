package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

/*

	分析:
	一、先把数字放在一个txt文档中
	二、用ioutil包的ReadFile函数读取出数字
	三、把读出的数据转换成int类型, 并装入slice
	四、以13为步长截取slice，并将其连乘，并装入切片中；注意index越界，所以slice的长度要减去步长
	五、将连乘获得的切片中的数字比较，取出最大值，即得结果


*/

func main() {

	start := time.Now().UnixNano()

	//第一步和第二部
	filename := "./numbers2.txt"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	//第三步
	var intSlice []int
	s := strings.Replace(string(bytes), "\n", "", -1)
	strSlice := strings.Split(s, "")
	for _, snum := range strSlice {
		inum, err := strconv.Atoi(snum)
		if err != nil {
			fmt.Println(err)
			return
		}
		intSlice = append(intSlice, inum)
	}

	//第四步
	var totalSlice []int
	l := len(intSlice)
	for i := 0; i < l-13; i++ {
		total := 1

		for j, num := range intSlice[i : i+13] {
			if num == 0 {
				i = i + j
				break
			}
			total *= num
			totalSlice = append(totalSlice, total)
			// fmt.Printf("i = %d, j = %d, num = %d, total = %d\n", i, j, num, total)
		}
	}

	//第五步
	var result int
	for _, totalNum := range totalSlice {
		if totalNum > result {
			result = totalNum
		}
	}
	fmt.Println(result)

	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000)
}
