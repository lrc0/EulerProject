package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().UnixNano()

	for a := 1; a < 333; a++ {
		for b := 1; b < 666; b++ {
			c := 1000 - a - b
			a2 := a * a
			b2 := b * b
			c2 := c * c
			if a2+b2 == c2 && a < b && b < c {
				fmt.Printf("a = %d, b = %d, c=%d\n", a, b, c)
				fmt.Println("abc = ", a*b*c)
			}
		}
	}

	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000)
}
