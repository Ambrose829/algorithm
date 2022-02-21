package binary_tree

import "math"

/**
  @Author: pika
  @Date: 2022/2/21 10:58 上午
  @Description: 二叉树
*/

/**
给定一个不重复的整数数组nums 。最大二叉树可以用下面的算法从nums 递归地构建:
创建一个根节点，其值为nums 中的最大值。
递归地在最大值左边的子数组前缀上构建左子树。
递归地在最大值 右边 的子数组后缀上构建右子树。
返回nums 构建的 最大二叉树 。
*/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var maxVal = math.MinInt32
	var index int
	//var root TreeNode
	//找到数组中最大值
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			index = i
		}
	}
	root := &TreeNode{
		Val: maxVal,
	}
	//递归调用构建左右子树
	root.Left = constructMaximumBinaryTree(nums[0:index])
	root.Right = constructMaximumBinaryTree(nums[index+1 : len(nums)-1])

	return root
}

/**
给定两个整数数组preorder 和 inorder，其中preorder 是二叉树的先序遍历， inorder是同一棵树的中序遍历，请构造二叉树并返回其根节点。
*/

func buildTree1(preorder []int, inorder []int) *TreeNode {
	return buildTreeByPreAndIn(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1)
}

func buildTreeByPreAndIn(preorder []int, preStart, preEnd int,
	inorder []int, inStart, inEnd int) *TreeNode {
	//base case
	if preStart > preEnd {
		return nil
	}

	//root节点为前序遍历的第一个值
	rootVal := preorder[preStart]
	//rootVal 在中序遍历数组中的位置
	var index int
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}
	//左子树的节点数
	leftLen := index - inStart

	root := &TreeNode{
		Val: rootVal,
	}
	//递归构造左右子树
	root.Left = buildTreeByPreAndIn(preorder, preStart+1, preStart+leftLen,
		inorder, inStart, index-1)

	root.Right = buildTreeByPreAndIn(preorder, preStart+leftLen+1, preEnd,
		inorder, index+1, inEnd)

	return root
}

/**
  给定两个整数数组inorder和postorder ，其中 norder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗二叉树。
*/

func buildTree2(inorder []int, postorder []int) *TreeNode {
	return buildTreeByInAndPost(inorder, 0, len(inorder)-1,
		postorder, 0, len(postorder)-1)
}

func buildTreeByInAndPost(inorder []int, inStart, inEnd int,
	postorder []int, postStart, postEnd int) *TreeNode {
	//base case
	if inStart > inEnd {
		return nil
	}
	rootVal := postorder[postEnd]
	var index int
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
			break
		}
	}

	leftlen := index - inStart

	root := &TreeNode{Val: rootVal}

	root.Left = buildTreeByInAndPost(inorder, inStart, index-1,
		postorder, postStart, postStart+leftlen-1)

	root.Right = buildTreeByInAndPost(inorder, index+1, inEnd,
		postorder, postStart+leftlen, postEnd-1)

	return root
}

/**
  给定两个整数数组，preorder和 postorder ，其中 preorder 是一个具有 无重复 值的二叉树的前序遍历，postorder 是同一棵树的后序遍历，重构并返回二叉树。
*/

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return buildTreeByPreAndPost(preorder, 0, len(preorder)-1,
		postorder, 0, len(postorder)-1)
}

func buildTreeByPreAndPost(preorder []int, preStart, preEnd int,
	postorder []int, postStart, postEnd int) *TreeNode {

	//base case
	if preStart > preEnd {
		return nil
	}

	if preStart == preEnd {
		return &TreeNode{Val: preorder[preStart]}
	}
	// root 节点对应的值就是前序遍历数组的第一个元素
	rootVal := preorder[preStart]
	//root.left就是前序遍历数组的第二个元素
	//通过前序遍历和后序遍历的关键是找出preorder和postorder的左右子树元素区间
	leftVal := preorder[preStart+1]
	// root.left 值 在后序遍历数组中的索引
	var index int
	for i := postStart; i <= postEnd; i++ {
		if postorder[i] == leftVal {
			index = i
			break
		}
	}

	//左子树元素个数，记得加个1
	leftLen := index - postStart + 1

	root := &TreeNode{Val: rootVal}

	root.Left = buildTreeByPreAndPost(preorder, preStart+1, preStart+leftLen,
		postorder, postStart, index)

	root.Right = buildTreeByPreAndPost(preorder, preStart+leftLen+1, preEnd,
		postorder, index+1, preEnd-1)

	return root
}
