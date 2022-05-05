package search

import (
	"math/rand"
	"time"
)

/**
  @Author: pika
  @Date: 2022/4/29 7:01 下午
  @Description: 随机加权选择
*/

/*
	给你一个 下标从 0 开始 的正整数数组w ，其中w[i] 代表第 i 个下标的权重。
	请你实现一个函数pickIndex，它可以 随机地 从范围 [0, w.length - 1] 内（含 0 和 w.length - 1）选出并返回一个下标。选取下标 i的 概率 为 w[i] / sum(w) 。
	例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3)= 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3)= 0.75（即，75%）。

*/

type Solution struct {
	preSum []int
}

func Constructor(w []int) Solution {

	preSum := make([]int, len(w))
	preSum[0] = w[0]
	for i := 1; i < len(w); i++ {
		preSum[i] = preSum[i-1] + w[i]
	}
	return Solution{preSum}
}

func (this *Solution) PickIndex() int {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(this.preSum[len(this.preSum)-1]) + 1
	return bsLeftBound(this.preSum, target)

}

func bsLeftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return left
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
