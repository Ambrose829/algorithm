package linked_list

/**
  @Author: pika
  @Date: 2022/1/26 11:20 上午
  @Description: 寻找单链表的倒数第 k 个节点
*/

/*
	使用快慢指针，指针间隔为k
*/

//返回列表的倒数第K个结点
func findKEnd(head ListNode, k int) ListNode {
	//快指针
	p1 := &head
	//p1先走k步
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}

	//慢指针
	p2 := &head

	//p1与p2 同时走n - k步
	for p1 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	//p2现在指向第n-k个节点
	return *p2
}

//单链表的中点
//快慢指针，快指针走两步=慢指针走一步

//判断链表是否包含环
func hasCycle(head ListNode) bool {
	//快慢指针初始化指向 head
	slowPtr := &head
	fastPtr := &head
	for fastPtr != nil && fastPtr.Next != nil {
		//慢指针走一步，快指针走两步
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		//快慢指针相遇，说明有环
		if slowPtr == fastPtr {
			return true
		}
	}
	//不包含环
	return false
}

//如何寻找链表中环的起点
//先找到相遇的点，然后再新增一个以头结点开始的同速指针
