package fibo

/**
  @Author: pika
  @Date: 2022/1/21 11:20 上午
  @Description: 斐波那契数列
*/


// 带备忘录的递归解法
func fibRec(n int) int {
	memo := make([]int, n + 1)
	return fibRecHelper(memo, n)
}

func fibRecHelper(memo []int, n int) int {
	// base case
	if n == 0 || n == 1 {
		return n
	}
	// 已经计算过了，不用再计算了
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fibRecHelper(memo, n - 1) + fibRecHelper(memo, n -2)
	return memo[n]
}


// dp数组的迭代解法
func fibDp(n int) int {
	if n == 0 {
		return 0
	}
	// dp table
	dp := make([]int, n + 1)
	// base case
	dp[0] = 0
	dp[1] = 1
	// 状态转移
	for i := 2; i <= n; i ++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
}
