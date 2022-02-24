package binary_tree

/**
  @Author: pika
  @Date: 2022/2/24 7:23 下午
  @Description: 二叉搜索树
*/

/**
BST特性：
	1、对于BST的每个节点node，左子树节点的值都比node的值要小，右子树节点的值都比node的值大
	2、对于BST的每个节点node，它的左侧子树和右侧子树都是BST
从算法的角度来看BST，它有一个重要的性质：BST的中序遍历结果是有序的（升序）
*/

/**
题目：二叉搜索树中第K小的元素
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

*/

func kthSmallest(root *TreeNode, k int) int {
	var res, num int
	var traverseFindTheK func(root *TreeNode, k int)
	traverseFindTheK = func(root *TreeNode, k int) {
		if root == nil {
			return
		}
		traverseFindTheK(root.Left, k)
		// 因为BST的中序遍历是升序的状态，判断逻辑
		// 判断逻辑
		num++
		if num == k {
			res = root.Val
			return
		}
		traverseFindTheK(root.Right, k)
	}
	traverseFindTheK(root, k)

	return res
}

/**
题目：把二叉搜索树转换为累加树
给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node的新值等于原树中大于或等于node.val的值之和。

思考：主要使用二叉搜索树的特性
*/
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var traverseBST2GST func(root *TreeNode)
	traverseBST2GST = func(root *TreeNode) {
		if root == nil {
			return
		}
		//先访问右子树，再访问左子树然后中序遍历就可以得到降序的数据
		traverseBST2GST(root.Right)
		sum += root.Val
		root.Val = sum
		traverseBST2GST(root.Left)
	}

	traverseBST2GST(root)

	return root
}
