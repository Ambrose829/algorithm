package coin_change

import "math"

/**
  @Author: pika
  @Date: 2022/1/21 3:15 下午
  @Description: 凑零钱问题
*/

/*
	题目：给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。
	算法的函数签名如下：
	```
	// coins 中是可选硬币面值，amount 是目标金额
	int coinChange(int[] coins, int amount)
	```
	比如说 k = 3，面值分别为 1，2，5，总金额 amount = 11。那么最少需要 3 枚硬币凑出，即 11 = 5 + 5 + 1。
 */

/*
	分析过程：
	明确base case：目标金额为0时返回0，因为不需要任何硬币就已经凑出目标金额amount。
	明确「状态」：愿问题和子问题中会变化的变量。由于硬币数量不限，硬币面额固定，只有目标面额不断的向base case接近，所以唯一的「状态」就是目标金额amount
	明确「选择」：导致「状态」产生变化的行为。目标金额变化的原因是在选择不同的硬币，每选择一枚硬币，就相当于减少了目标金额。
	明确dp函数/数组定义：

 */

//递归
func coinChangeRec(coins []int, amount int) int {
	return coinChangeRecHelper(coins, amount)


}

func coinChangeRecHelper(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	res := math.MaxInt32
	for _, coin := range coins {
		//计算子问题结果
		subProblem := coinChangeRecHelper(coins, amount - coin)
		//子问题无则跳过
		if subProblem == -1 {
			continue
		}
		//在子问题中寻找最优解，然后加1
		if subProblem < res {
			res = subProblem + 1
		}
	}
	if res == math.MaxInt32 {
		res = -1
	}
	return res
}



//备忘录递归
var memo []int

func coinChangeRecMem(coins []int, amount int) int {
	//初始化备忘录（dp数组）
	memo =  make([]int, amount + 1)
	return coinChangeRecMemHelper(coins, amount)
}

func coinChangeRecMemHelper(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] != 0 {
		return memo[amount]
	}

	res := math.MaxInt32
	for _, coin := range coins {
		//计算子问题结果
		subProblem := coinChangeRecMemHelper(coins, amount - coin)
		//子问题无则跳过
		if subProblem == -1 {
			continue
		}
		//在子问题中寻找最优解，然后加1
		if subProblem < res {
			res = subProblem + 1
		}
	}
	if res == math.MaxInt32 {
		res = -1
		memo[amount] = res
	}
	return res
}



//dp数组的迭代解法
func coinChargeDp(coins []int, amount int) int {
	//初始化数组
	dp := make([]int, amount + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	//base case
	dp[0] = 0
	//外层for循环遍历所有状态的取值
	for i := 0; i < len(dp); i ++ {
		//内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			//子问题无解跳过
			if i - coin < 0 {
				continue
			}
			if dp[i - coin] < dp[i] {
				dp[i] = dp[i - coin] + 1
			}
		}
	}
	//无解
	if dp[amount] == amount + 1 {
		dp[amount] = -1
	}
	return dp[amount]
}
