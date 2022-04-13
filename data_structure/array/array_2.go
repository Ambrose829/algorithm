package array

/**
  @Author: pika
  @Date: 2022/4/11 9:46 下午
  @Description: 差分数组
				差分数组的主要适用场景是频繁对原始数组的某个区间的元素进行增减。
*/

// 差分数组工具类

//差分数组
var diff []int

// Difference 输入初始数组，生成差分数组
func Difference(nums []int) {
	if len(nums) < 1 {
		return
	}
	diff = make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
}

//

func Increment(i, j, val int) {
	diff[i] += val
	if j+1 < len(diff) {
		diff[j+1] -= val
	}
}

func Result() []int {
	var res []int
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = diff[i] + res[i-1]
	}

	return res
}

/*
	假设你有一个长度为n的数组初识情况下所有数字均为0，你将会被给出k个更新的操作。
	其中，每个操作会被表示为一个三元组：[startIndex, endIndex, inc], 你需要将子数组A[startIndex ... endIndex](包括 startIndex和endIndex)增加inc。

	请你返回k次操作后的数组
*/

func getModifiedArray(length int, updates [][]int) []int {
	nums := make([]int, length)
	Difference(nums)
	for _, update := range updates {
		i := update[0]
		j := update[1]
		val := update[2]
		Increment(i, j, val)
	}
	return Result()
}

/*
	这里有n个航班，它们分别从 1 到 n 进行编号。
	有一份航班预订表bookings ，表中第i条预订记录bookings[i] = [firsti, lasti, seatsi]意味着在从 firsti到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi个座位。
	请你返回一个长度为 n 的数组answer，里面的元素是每个航班预定的座位总数。
*/

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)

	//差分数组
	var diff []int

	//生成差分数组
	var difference func(nums []int)

	difference = func(nums []int) {
		if len(nums) < 1 {
			return
		}
		diff = make([]int, len(nums))
		diff[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			diff[i] = nums[i] - nums[i-1]
		}
	}

	//差分数组更新
	var increment func(i, j, val int)

	increment = func(i, j, val int) {
		diff[i] += val
		if j+1 < len(diff) {
			diff[j+1] -= val
		}
	}

	//根据差分数组返回值
	var result func() []int
	result = func() []int {
		res := make([]int, len(diff))
		res[0] = diff[0]
		for i := 1; i < len(diff); i++ {
			res[i] = diff[i] + res[i-1]
		}
		return res
	}

	difference(nums)

	for _, booking := range bookings {
		increment(booking[0]-1, booking[1]-1, booking[2])
	}

	res := result()
	return res

}

/*
	车上最初有capacity个空座位。车只能向一个方向行驶（也就是说，不允许掉头或改变方向）

	给定整数capacity和一个数组 trips , trip[i] = [numPassengersi, fromi, toi]表示第 i 次旅行有numPassengersi乘客，接他们和放他们的位置分别是fromi和toi。这些位置是从汽车的初始位置向东的公里数。

	当且仅当你可以在所有给定的行程中接送所有乘客时，返回true，否则请返回 false。
	。
*/

func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)

	//差分数组
	var diff []int

	//生成差分数组
	var difference func(nums []int)

	difference = func(nums []int) {
		if len(nums) < 1 {
			return
		}
		diff = make([]int, len(nums))
		diff[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			diff[i] = nums[i] - nums[i-1]
		}
	}

	//差分数组更新
	var increment func(i, j, val int)

	increment = func(i, j, val int) {
		diff[i] += val
		if j+1 < len(diff) {
			diff[j+1] -= val
		}
	}

	//根据差分数组返回值
	var result func() []int
	result = func() []int {
		res := make([]int, len(diff))
		res[0] = diff[0]
		for i := 1; i < len(diff); i++ {
			res[i] = diff[i] + res[i-1]
		}
		return res
	}

	difference(nums)

	for _, trip := range trips {
		increment(trip[1], trip[2]-1, trip[0])
	}

	res := result()
	for _, re := range res {
		if re > capacity {
			return false
		}
	}

	return true
}
