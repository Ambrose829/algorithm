package array

import (
	"math"
	"strings"
)

/**
  @Author: pika
  @Date: 2022/4/19 4:11 下午
  @Description: 滑动窗口
*/

func template(arrays []int) {
	//左右指针,左闭右开
	left, right := 0, 0

	// 窗口
	var windows []int

	// 窗口缩小条件
	var needShrink bool

	for right < len(arrays) {

		//增大窗口
		windows = append(windows, arrays[right])
		right++
		for needShrink {
			// 缩小窗口
			windows = windows[left:]
			left++
		}
	}
}

/*
	给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
	注意：

	对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
	如果 s 中存在这样的子串，我们保证它是唯一的答案。
*/

func minWindow1(s string, t string) string {
	// 初始化窗口和需要判断的值
	var needs, windows map[string]int

	// 初始化需要校验的子串和滑动窗口
	needs = make(map[string]int, len(t))
	windows = make(map[string]int, 1)
	ss := strings.Split(s, "")
	ts := strings.Split(t, "")

	//填充需要验证覆盖的信息
	for _, s2 := range ts {
		_, isExist := needs[s2]
		if isExist {
			needs[s2] += 1
		} else {
			needs[s2] = 1
		}
	}

	//滑动窗口的左右边界
	left, right := 0, 0

	// 滑动窗口中对应的有需要一样的字符个数
	valid := 0

	//最小覆盖字串的起始值和长度
	start, length := 0, math.MaxInt32

	for right < len(ss) {
		// 放入滑动窗口中的字符
		c := ss[right]

		//扩大窗口
		right++

		// 进行滑动窗口内一系列数据的更新
		// 是滑动窗口中的要验证的字段
		_, needExist := needs[c]
		if needExist {
			// 将字符添加进滑动窗口
			_, isExist := windows[c]
			if isExist {
				windows[c] += 1
			} else {
				windows[c] = 1
			}

			// 以字符为维度计算限制值
			if needs[c] == windows[c] {
				valid++
			}
		}

		// 滑动窗口中已经覆盖了需要校验的子串
		for valid == len(needs) {

			// 记录并更新最小覆盖子串
			if right-left < length {
				//起始位置
				start = left
				//子串长度
				length = right - left
			}
			// d是即将删除的字符
			d := ss[left]

			//缩小滑动窗口
			left++
			_, needExist2 := needs[d]
			if needExist2 {
				if needs[d] == windows[d] {
					valid--
				}
				windows[d]--
			}
		}

	}
	if length == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+length]
	}

}

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

/*
	给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。
	换句话说，s1 的排列之一是 s2 的 子串 。

*/

func checkInclusion(s1 string, s2 string) bool {
	//滑动窗口的左右边界
	left, right := 0, 0

	//滑动窗口和要校验的窗口
	var needs, windows map[string]int
	needs = make(map[string]int, len(s1))
	windows = make(map[string]int, 1)

	//初始化/数据填充
	s1s := strings.Split(s1, "")
	s2s := strings.Split(s2, "")

	for _, s := range s1s {
		_, f1 := needs[s]
		if f1 {
			needs[s] += 1
		} else {
			needs[s] = 1
		}
	}

	var valid int

	// 数据处理
	for right < len(s2s) {
		var c = s2s[right]
		_, f2 := needs[c]
		if f2 {
			_, f3 := windows[c]
			if f3 {
				windows[c] += 1
			} else {
				windows[c] = 1
			}
			if windows[c] == needs[c] {
				valid++
			}
		}

		right++

		// 缩小窗口,当滑动窗口的大小与需要校验的排列长度不一致时
		for right-left >= len(s1s) {
			if valid == len(needs) {
				return true
			}
			var d = s2s[left]
			left++
			_, f4 := needs[d]
			if f4 {
				if windows[d] == needs[d] {
					valid--
				}
				windows[d]--
			}

		}

	}

	return false
}

/*
	给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
	异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

*/
func findAnagrams(s string, p string) []int {
	var res []int

	// 滑动窗口左右边界
	left, right := 0, 0

	// 定义窗口和校验值
	var windows, needs map[string]int
	needs = make(map[string]int, len(p))
	windows = make(map[string]int, 1)
	ps := strings.Split(p, "")
	for _, s2 := range ps {
		needs[s2]++
	}
	ss := strings.Split(s, "")

	var valid int
	for right < len(ss) {
		c := ss[right]
		_, ok1 := needs[c]
		if ok1 {
			windows[c]++
			if needs[c] == windows[c] {
				valid++
			}
		}

		right++

		for right-left >= len(p) {

			if valid == len(needs) {
				res = append(res, left)
			}

			d := ss[left]
			left++
			_, f3 := needs[d]
			if f3 {
				if needs[d] == windows[d] {
					valid--
				}
				windows[d]--
			}

		}
	}
	return res
}

/*

	给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度

*/

func lengthOfLongestSubstring(s string) int {
	var resLen int

	// 滑动窗口左右边界
	left, right := 0, 0

	// 定义窗口和校验值
	var windows map[string]int
	windows = make(map[string]int, 1)

	ss := strings.Split(s, "")

	for right < len(s) {
		c := ss[right]
		windows[c]++

		right++

		for windows[c] > 1 {

			d := ss[left]
			left++

			if windows[d] > 1 {
				windows[d]--
			} else {
				delete(windows, d)
			}

		}

		if resLen < len(windows) {
			resLen = len(windows)
		}

	}

	return resLen
}
