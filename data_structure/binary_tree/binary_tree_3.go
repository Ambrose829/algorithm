package binary_tree

import "fmt"

/**
  @Author: pika
  @Date: 2022/2/22 9:28 下午
  @Description: 二叉树遍历
*/

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	//存放子树集
	var subTreeMap map[string]int
	//返回结果集
	var res []*TreeNode
	subTreeMap = make(map[string]int, 1)
	var traverseSub func(root *TreeNode) string
	traverseSub = func(root *TreeNode) string {
		if root == nil {
			return "*"
		}
		left := traverseSub(root.Left)
		right := traverseSub(root.Right)
		s := fmt.Sprintf("%d %s %s", root.Val, left, right)
		v, isExist := subTreeMap[s]
		if isExist && v == 1 {
			res = append(res, root)
			subTreeMap[s] = v + 1
		} else {
			if s != "*" {
				subTreeMap[s] = v + 1
			}
		}

		return s
	}
	traverseSub(root)
	return res
}

//    0
//  0   0
// 0   0
//0 0 0 0
//
//[0,0,0,0,null,0,null,0,0,0,0]
//输出：
//[[0],[0,0,0]]
//预期结果：
//[[0,0,null,0,0],[0,0,0],[0]]
//
//
//   2
// 1   11
//11  1
