package binary_tree

/**
  @Author: pika
  @Date: 2022/2/15 8:07 下午
  @Description: 二叉树基础
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	//前序位置
	traverse(root.Left)
	//中序位置
	traverse(root.Right)
	//后序位置
}

/**
	二叉树的直径：
     给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//错误解法
var maxDiameter int

func diameterOfBinaryTree1(root *TreeNode) int {
	maxDepth(root)
	return maxDiameter
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	myDiameter := leftMax + rightMax
	maxDiameter = max(maxDiameter, myDiameter)
	return max(leftMax, rightMax) + 1
}

//正确解法
func diameterOfBinaryTree2(root *TreeNode) int {
	// 最大直径=左子树高度+右子树高度
	maxDia := 0
	if root == nil {
		return maxDia
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lHeight := dfs(node.Left)
		rHeight := dfs(node.Right)
		// maxDia即为该二叉树的最大直径
		maxDia = max(maxDia, lHeight+rHeight)
		// 该节点为根的二叉树的最大高度，也就是1+max(lh+rh)
		return 1 + max(lHeight, rHeight)
	}
	dfs(root)
	return maxDia
}
