package search

import "sort"

/**
  @Author: pika
  @Date: 2022/5/19 6:39 下午
  @Description: 优势决策， 田忌赛马
*/

/*
	给定两个大小相等的数组nums1和nums2，nums1相对于 nums的优势可以用满足nums1[i] > nums2[i]的索引 i的数目来描述。

	返回 nums1的任意排列，使其相对于 nums2的优势最大化。

	示例 1：

	输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
	输出：[2,11,7,15]
	示例 2：

	输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
	输出：[24,32,8,12]


*/

type pair struct {
	x int
	y int
}

func advantageCount1(nums1 []int, nums2 []int) []int {

	// 将我方数组升序排序
	sort.Ints(nums1)

	// 构建对方数组的信息
	n2 := make([]pair, len(nums2))

	for i := range nums2 {
		n2[i].x = i
		n2[i].y = nums2[i]
	}

	sort.Slice(n2, func(i, j int) bool {
		return n2[i].y < n2[j].y
	})

	// 返回值
	res := make([]int, len(nums1))

	// 左右区间
	l, r := 0, len(nums1)-1

	for i := 0; i < len(nums1); i++ {
		if nums1[i] > n2[l].y {
			res[n2[l].x] = nums1[i]
			l++
		} else {
			res[n2[r].x] = nums1[i]
			r--
		}
	}

	return res
}

func advantageCount(nums1 []int, nums2 []int) []int {

	// 将我方数组升序排序
	sort.Ints(nums1)

	// 构建对方数组的信息
	n2 := make([]pair, len(nums2))

	for i := range nums2 {
		n2[i].x = i
		n2[i].y = nums2[i]
	}

	sort.Slice(n2, func(i, j int) bool {
		return n2[i].y > n2[j].y
	})

	// 返回值
	res := make([]int, len(nums1))

	// 左右区间
	l, r := 0, len(nums1)-1

	for i := range n2 {
		if nums1[r] > n2[i].y {
			res[n2[i].x] = nums1[r]
			r--
		} else {
			res[n2[i].x] = nums1[l]
			l++
		}
	}

	return res
}
