package array

/**
  @Author: pika
  @Date: 2022/3/24 4:58 下午
  @Description: 前缀和数组
*/

/*
	给定一个整数数组 nums，处理以下类型的多个查询:

	计算索引left和right（包含 left 和 right）之间的 nums 元素的 和 ，其中left <= right
	实现 NumArray 类：

	NumArray(int[] nums) 使用数组 nums 初始化对象
	int sumRange(int i, int j) 返回数组 nums中索引left和right之间的元素的 总和 ，包含left和right两点（也就是nums[left] + nums[left + 1] + ... + nums[right])

*/

// 使用前缀和解法

type NumArray struct {
	pnums []int
}

func Constructor1(nums []int) NumArray {
	var pnums = make([]int, len(nums)+1)
	for i := 1; i < len(pnums); i++ {
		pnums[i] = pnums[i-1] + nums[i-1]
	}
	return NumArray{pnums: pnums}
}

func (this *NumArray) SumRange(left int, right int) int {

	return this.pnums[right+1] - this.pnums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

/*
	给定一个二维矩阵 matrix，以下类型的多个请求：

	计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1,col1) ，右下角 为 (row2,col2) 。
	实现 NumMatrix 类：

	NumMatrix(int[][] matrix)给定整数矩阵 matrix 进行初始化
	int sumRegion(int row1, int col1, int row2, int col2)返回
	左上角 (row1,col1)、右下角(row2,col2) 所描述的子矩阵的元素 总和 。

*/

type NumMatrix struct {
	preMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {

	preMatrix := make([][]int, len(matrix)+1)
	for i := range preMatrix {
		preMatrix[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			preMatrix[i][j] = matrix[i-1][j-1] + preMatrix[i-1][j] + preMatrix[i][j-1] - preMatrix[i-1][j-1]
		}
	}
	return NumMatrix{preMatrix: preMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	//注意这里要减掉的部分是 [row1 + 1 - 1][col2 + 1] 和 [row2 + 1][col1 + 1 - 1]
	return this.preMatrix[row2+1][col2+1] - this.preMatrix[row1][col2+1] - this.preMatrix[row2+1][col1] + this.preMatrix[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

/*
	给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
*/

func subarraySum(nums []int, k int) int {

	preNum := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		preNum[i+1] = preNum[i] + nums[i]
	}
	var res int
	for i := 0; i < len(preNum); i++ {
		for j := 0; j < i; j++ {
			if preNum[i]-preNum[j] == k {
				res++
			}
		}
	}
	return res

}
