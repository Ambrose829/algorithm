package linked_list

/**
  @Author: pika
  @Date: 2022/1/24 7:36 下午
  @Description: 合并两个排序列表
*/

/*
	题目：将两个升序链表合并为一个新的升序链表返回。新链表升通过拼接给定的两个链表的所有节点组成的。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := new(ListNode)
	p := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}

	return res.Next
}
