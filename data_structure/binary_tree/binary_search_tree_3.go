package binary_tree

/**
  @Author: pika
  @Date: 2022/2/28 10:42 上午
  @Description: 二叉搜索树
*/

//如何计算所有合法的BST

/**
题目：不同的二叉搜索树
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
*/

// 单纯递归解法
func numTrees1(n int) int {
	return countNumTrees1(1, n)
}

func countNumTrees1(low, high int) int {
	if low > high {
		return 1
	}
	var res int
	for i := low; i <= high; i++ {
		left := countNumTrees1(low, i-1)
		right := countNumTrees1(i+1, high)
		res += left * right
	}
	return res
}

// 带备忘录的递归解法
var memo [][]int

func numTrees(n int) int {
	memo = make([][]int, n)
	for i, _ := range memo {
		memo[i] = make([]int, n)
	}
	return countNumTrees(1, n)
}

func countNumTrees(low, high int) int {
	if low > high {
		return 1
	}
	if memo[low][high] != 0 {
		return memo[low][high]
	}
	var res int
	for rt := low; rt <= high; rt++ {
		left := countNumTrees(low, rt-1)
		right := countNumTrees(rt+1, high)
		res += left * right
	}
	memo[low][high] = res

	return res
}

// 动态规划解法
func numTreesDp(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] += dp[j] * dp[i-j]
		}
	}
	return dp[n]
}

/**
题目：不同的二叉搜索树 II
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
*/

func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return nil
	}
	return generateBST(1, n)
}

func generateBST(lo, hi int) []*TreeNode {
	var bsts []*TreeNode

	//base case
	if lo > hi {
		bsts = append(bsts, nil)
		return bsts
	}
	for i := lo; i <= hi; i++ {
		leftTrees := generateBST(lo, i-1)
		rightTrees := generateBST(i+1, hi)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				bsts = append(bsts, root)
			}
		}
	}
	return bsts
}
