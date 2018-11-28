package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"time"
)

func main() {

	start := time.Now().UnixNano()

	filename := "./numbers.txt"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(bytes)
	slice := strings.Split(str, "\n")

	i := new(big.Int)
	y := new(big.Int)

	for _, numStr := range slice {
		//用math中的big包将numStr转换成*Int类型
		x, ok := i.SetString(numStr, 10)
		if ok {
			//将转换之后的数连加起来，注意这个y,因为Add函数返回的值是保存在y中的
			y = y.Add(y, x)
		} else {
			fmt.Println(numStr)
		}
	}
	resultNum := y.String()
	fmt.Println(resultNum)
	fmt.Println("result = ", resultNum[:10])

	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000, "微秒")

}
