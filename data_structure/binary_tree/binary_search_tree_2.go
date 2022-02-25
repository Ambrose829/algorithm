package binary_tree

/**
  @Author: pika
  @Date: 2022/2/25 4:16 下午
  @Description: 二叉搜索树
*/

/**
题目：二叉搜索树中的搜索
你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 null 。
*/

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	var res *TreeNode
	if val < root.Val {
		res = searchBST(root.Left, val)
	} else {
		res = searchBST(root.Right, val)
	}
	return res
}

/**
题目：验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
*/
func isValidBST(root *TreeNode) bool {
	return judgeBST(root, nil, nil)

}

func judgeBST(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}

	return judgeBST(root.Left, min, root) && judgeBST(root.Right, root, max)
}

/**
题目：二叉搜索树中的插入操作
给定二叉搜索树（BST）的根节点root和要插入树中的值value，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。
注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
*/

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	return traverseInsertBST(root, val)
}

func traverseInsertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val >= val {
		root.Left = traverseInsertBST(root.Left, val)
	}

	if root.Val < val {
		root.Right = traverseInsertBST(root.Right, val)
	}
	return root
}

/**
题目：删除二叉搜索树中的节点
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的key对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
一般来说，删除节点可分为两个步骤：
	首先找到需要删除的节点；
	如果找到了，删除它。
*/

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key == root.Val {
		//删除节点
		//左节点为空
		if root.Left == nil {
			return root.Right
		}
		//右节点为空
		if root.Right == nil {
			return root.Left
		}

		minNode := findMin(root.Right)
		root.Right = deleteNode(root.Right, minNode.Val)
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode

	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func findMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}
