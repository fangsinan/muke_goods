package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10000; i++ {
		go func(i int) {
			time.Sleep(2 * time.Second)
			fmt.Println("一个循环", i)
		}(i)
	}

	for i := 0; i < 10000; i++ {
		go func(i int) {
			time.Sleep(2 * time.Second)
			fmt.Println("两个循环", i)
		}(i)
	}

	for {
	}
}
