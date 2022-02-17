package linked_list

/**
  @Author: pika
  @Date: 2022/2/17 2:34 下午
  @Description: 判断回文链表
*/

/**
1、利用栈原理，使用链表的后序遍历，实现整个链表的倒序操作，与正序对比
2、使用双指针，以中间为刻度，将后序链表反转并与前半部分对比，确定是否为回文链表
*/

func isPalindrome(head *ListNode) bool {
	left := head
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}

	right := reverseLinkedList(slow)
	for right != nil {
		if right.Val != left.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}
