package coin_change

import (
	"fmt"
	"testing"
)

/**
  @Author: pika
  @Date: 2022/1/24 5:00 下午
  @Description:
*/

func Test_coinChangeRec(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 13
	num := coinChangeRec(coins, amount)
	fmt.Println(num)
}


func Test_coinChangeRecMem(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 13
	num := coinChangeRecMem(coins, amount)
	fmt.Println(num)
}

func Test_coinChargeDp(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 13
	num := coinChargeDp(coins, amount)
	fmt.Println(num)
}
