package binary_tree

import (
	"bytes"
	"math"
	"strconv"
	"strings"
)

/**
  @Author: pika
  @Date: 2022/3/1 11:27 上午
  @Description: 二叉搜索树补充
*/

/**
题目：二叉搜索子树的最大键值和
给你一棵以 root 为根的 二叉树 ，请你返回 任意 二叉搜索子树的最大键值和。
*/

func maxSumBST(root *TreeNode) int {
	var maxSum int

	var traverseMaxSumBST func(root *TreeNode) [4]int

	// [4]int{'是否为BST：0否；1是', 最小值, 最大值, 和}
	traverseMaxSumBST = func(root *TreeNode) [4]int {
		res := [4]int{0, 0, 0, 0}
		if root == nil {
			return [4]int{1, math.MaxInt32, math.MinInt32, 0}
		}
		leftRes := traverseMaxSumBST(root.Left)
		rightRes := traverseMaxSumBST(root.Right)
		if leftRes[0] == 1 && rightRes[0] == 1 &&
			root.Val > leftRes[2] && root.Val < rightRes[1] {
			res[0] = 1
			res[1] = min(leftRes[1], root.Val)
			res[2] = max(rightRes[2], root.Val)
			res[3] = leftRes[3] + rightRes[3] + root.Val
			maxSum = max(res[3], maxSum)
		}
		return res
	}

	traverseMaxSumBST(root)
	return maxSum
}

/**
题目：二叉树的序列化与反序列化
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

*/

const (
	SEP = ","
	NIL = "*"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var buffer bytes.Buffer
	recSerialize(root, &buffer)
	return buffer.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var strSlice []string
	strSlice = strings.Split(data, ",")
	var recDeserialize func() *TreeNode
	recDeserialize = func() *TreeNode {

		if strSlice[0] == NIL {
			strSlice = strSlice[1:]
			return nil
		}
		rootVal, _ := strconv.Atoi(strSlice[0])

		root := &TreeNode{Val: rootVal}
		strSlice = strSlice[1:]
		root.Left = recDeserialize()
		root.Right = recDeserialize()

		return root
	}

	return recDeserialize()
}

func recSerialize(root *TreeNode, buffer *bytes.Buffer) {
	if root == nil {
		buffer.WriteString(NIL)
		buffer.WriteString(SEP)
		return

	}
	buffer.WriteString(strconv.Itoa(root.Val))
	buffer.WriteString(SEP)
	recSerialize(root.Left, buffer)
	recSerialize(root.Right, buffer)

}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
