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

/*
	珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有piles[i]根香蕉。警卫已经离开了，将在 h 小时后回来。
	珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
	珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
	返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。


	输入：piles = [3,6,7,11], h = 8
	输出：4

	输入：piles = [30,11,23,4,20], h = 5
	输出：30


*/

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1000000000

	for left <= right {
		mid := left + (right-left)/2
		if f(piles, mid) == h {
			right = mid - 1
		} else if f(piles, mid) < h {
			right = mid - 1
		} else if f(piles, mid) > h {
			left = mid + 1
		}
	}

	return left
}

func f(piles []int, x int) int {
	var h int
	for _, pile := range piles {
		h += pile / x
		if pile%x > 0 {
			h++
		}
	}
	return h
}

/*
	传送带上的包裹必须在 days 天内从一个港口运送到另一个港口。
	传送带上的第 i个包裹的重量为weights[i]。每一天，我们都会按给出重量（weights）的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。
	返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。


	输入：weights = [1,2,3,4,5,6,7,8,9,10], days = 5
	输出：15
	解释：
	船舶最低载重 15 就能够在 5 天内送达所有包裹，如下所示：
	第 1 天：1, 2, 3, 4, 5
	第 2 天：6, 7
	第 3 天：8
	第 4 天：9
	第 5 天：10

	请注意，货物必须按照给定的顺序装运，因此使用载重能力为 14 的船舶并将包装分成 (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) 是不允许的。


*/

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, weight := range weights {
		right += weight
		if left < weight {
			left = weight
		}
	}

	for left <= right {
		mid := left + (right-left)/2
		if ship(weights, mid) == days {
			right = mid - 1
		} else if ship(weights, mid) < days {
			right = mid - 1
		} else if ship(weights, mid) > days {
			left = mid + 1
		}
	}
	return left
}

func ship(weights []int, w int) int {
	var d int
	var sum int
	for _, weight := range weights {
		sum += weight
		if sum > w {
			d++
			sum = weight
		}
	}
	// 无论如何最后都要加一天， 如果没有超载 d++不触发，如果超载还需要再送一次
	return d + 1
}
