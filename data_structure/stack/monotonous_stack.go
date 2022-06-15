package stack

import (
	"fmt"
)

/**
  @Author: pika
  @Date: 2022/6/14 8:44 下午
  @Description:
*/

/*
	给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
	示例 1：
	输入：s = "bcabc"
	输出："abc"

	示例 2：
	输入：s = "cbacdcbc"
	输出："acdb"

*/

func removeDuplicateLetters1(s string) string {
	countN := [256]int{}
	for _, c := range s {
		countN[c]++
	}

	var stackS Stack
	instack := [256]bool{}

	for _, c := range s {
		countN[c]--
		if instack[c] {
			continue
		}
		for !stackS.isEmpty() && stackS.peek() > c {
			if countN[stackS.peek()] == 0 {
				break
			}
			instack[stackS.pop()] = false
		}
		stackS.push(c)
		instack[c] = true
	}

	var res string
	for !stackS.isEmpty() {
		fmt.Sprintf("%c%s", stackS.pop(), res)
	}

	return res
}

type Stack []int32

// 入栈
func (s *Stack) push(a int32) {
	*s = append(*s, a)
}

// 出栈
func (s *Stack) pop() int32 {
	a := *s
	defer func() {
		*s = a[:len(a)-1]
	}()
	return a[len(a)-1]
}

func (s *Stack) peek() int32 {
	a := *s
	return a[len(a)-1]
}

func (s *Stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	} else {
		return false
	}
}

func removeDuplicateLetters(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}

	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		left[ch-'a']--
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'

				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
	}
	return string(stack)

}
