package array

/**
  @Author: pika
  @Date: 2022/4/13 4:47 下午
  @Description: 双指针
*/

/*
	给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。
	由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么nums的前 k 个元素应该保存最终结果。
	将最终结果插入nums 的前 k 个位置后返回 k 。
	不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
	判题标准:
	系统会用下面的代码来测试你的题解:
	int[] nums = [...]; // 输入数组
	int[] expectedNums = [...]; // 长度正确的期望答案

	int k = removeDuplicates(nums); // 调用

	assert k == expectedNums.length;
	for (int i = 0; i < k; i++) {
	    assert nums[i] == expectedNums[i];
	}
	如果所有断言都通过，那么您的题解将被 通过。

*/

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1

}

/*
	给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。

*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	slow, fast := head, head
	if head == nil || head.Next == nil {
		return head
	}

	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next

	}
	slow.Next = nil
	return head
}

/*
	给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
	不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
	元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
	说明:
	为什么返回数值是整数，但输出的答案是数组呢?
	请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
	你可以想象内部操作如下:


*/

func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}
	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

/*
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
	请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 滑动窗口解法(自己方式 无第一种效率高)
func moveZeroes2(nums []int) {
	left, right := 0, 0
	for left < len(nums) {
		for nums[left] == 0 && right < len(nums) {
			if nums[right] != 0 && right > left {
				nums[left], nums[right] = nums[right], nums[left]
			}
			right++
		}
		left++
	}
}

/*

	给你一个下标从 1 开始的整数数组numbers ，该数组已按 非递减顺序排列 ，请你从数组中找出满足相加之和等于目标数target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
	以长度为 2 的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。
	你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
	你所设计的解决方案必须只使用常量级的额外空间。

*/

func twoSum(numbers []int, target int) []int {

	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}

	return []int{-1, -1}
}

/*
	编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
	不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

*/
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

/*
	给你一个字符串 s，找到 s 中最长的回文子串。
*/

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		ps1 := fPalindrome(s, i, i)
		ps2 := fPalindrome(s, i, i+1)
		if len(res) < len(ps1) {
			res = ps1
		}
		if len(res) < len(ps2) {
			res = ps2
		}

	}
	return res
}

func fPalindrome(s string, l, r int) string {
	ss := []byte(s)
	for l >= 0 && r < len(ss) && ss[l] == ss[r] {
		l--
		r++
	}
	return string(ss[l+1 : r])
}
