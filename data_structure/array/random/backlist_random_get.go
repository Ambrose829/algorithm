package random

import (
	"math/rand"
	"time"
)

/**
  @Author: pika
  @Date: 2022/6/9 4:42 下午
  @Description: 黑名单中的随机数
*/

/*
	给定一个整数 n 和一个 无重复 黑名单整数数组blacklist。
	设计一种算法，从 [0, n - 1] 范围内的任意整数中选取一个未加入黑名单blacklist的整数。
	任何在上述范围内且不在黑名单blacklist中的整数都应该有 同等的可能性 被返回。

	优化你的算法，使它最小化调用语言 内置 随机函数的次数。

	实现Solution类:

	Solution(int n, int[] blacklist)初始化整数 n 和被加入黑名单blacklist的整数
	int pick()返回一个范围为 [0, n - 1] 且不在黑名单blacklist 中的随机整数


	输入
	["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
	[[7, [2, 3, 5]], [], [], [], [], [], [], []]
	输出
	[null, 0, 4, 1, 6, 1, 0, 4]

	解释
	Solution solution = new Solution(7, [2, 3, 5]);
	solution.pick(); // 返回0，任何[0,1,4,6]的整数都可以。注意，对于每一个pick的调用，
	                 // 0、1、4和6的返回概率必须相等(即概率为1/4)。
	solution.pick(); // 返回 4
	solution.pick(); // 返回 1
	solution.pick(); // 返回 6
	solution.pick(); // 返回 1
	solution.pick(); // 返回 0
	solution.pick(); // 返回 4



*/

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */

type Solution struct {
	elen    int
	mapping map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	elen := n - len(blacklist)
	blacklistM := make(map[int]int, len(blacklist))
	for i := 0; i < len(blacklist); i++ {
		blacklistM[blacklist[i]] = 666
	}

	last := n - 1
	for i := 0; i < len(blacklist); i++ {
		if blacklist[i] >= elen {
			continue
		}

		for {
			_, ok := blacklistM[last]
			if ok {
				last--
			} else {
				blacklistM[blacklist[i]] = last
				last--
				break
			}
		}
	}

	return Solution{
		elen:    elen,
		mapping: blacklistM,
	}
}

func (this *Solution) Pick() int {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(this.elen)
	if i, ok := this.mapping[idx]; ok {
		return i
	}
	return idx
}
