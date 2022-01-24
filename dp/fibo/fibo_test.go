package fibo

import (
	"fmt"
	"testing"
)

/**
  @Author: pika
  @Date: 2022/1/21 2:23 下午
  @Description:
*/

//测试fib
func Test_fibRec(t *testing.T) {
	fib7 := fibRec(7)
	fib8 := fibRec(8)
	fib9 := fibRec(9)
	fmt.Println(fib7, fib8, fib9)
}


func Test_fibDp(t *testing.T) {
	fib7 := fibDp(7)
	fib8 := fibDp(8)
	fib9 := fibDp(9)
	fmt.Println(fib7, fib8, fib9)
}