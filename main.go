package main

import (
	"errors"
	"fmt"
	"time"
)

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
func GoOne() int {
	return 1
}

func GoTwo() int {

	return 2

}

func GoThree() int {

	return 3

}

func Sum(da []interface{}) error {

	sum := 0
	for i, _ := range da {
		f, ok := da[i].(func() int)
		if ok {
			sum += f()
		} else {
			return errors.New("sdf")
		}

	}
	fmt.Println("sum:", sum)
	return nil

}

func main() {
	a := []interface{}{GoOne, GoTwo, GoThree}
	go Sum(a)
	time.Sleep(100 * time.Second)

}
