package search

/**
  @Author: pika
  @Date: 2022/4/29 3:16 下午
  @Description: 二分查找
*/

/*
	给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

*/

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

// 二分查找左边界，左闭右开
func binarySearchLeftBd(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}

// 二分查找左边界统一，左闭右闭
func binarySearchLeftBdGeneral(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 二分查找右边界，左闭右开
func binarySearchRightBd(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}

	return left - 1
}

// 二分查找右边界统一，左闭右闭
func binarySearchRightBdGeneral(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}

	return right
}

/*
	给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
	如果数组中不存在目标值 target，返回[-1, -1]。
	进阶：
	你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？

*/

func searchRange(nums []int, target int) []int {

	leftBound := binarySearchLeftBdGeneral(nums, target)
	rightBound := binarySearchRightBdGeneral(nums, target)

	return []int{leftBound, rightBound}
}
