package linked_list

import "container/heap"

/**
  @Author: pika
  @Date: 2022/1/24 8:05 下午
  @Description: 合并k个有序链表
*/

/**
题目：给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/

/**
问题分析：合并n个有序链表的逻辑类似合并两个有序链表，难点在于，如何快速得到n个节点中的最小节点，接到结果链表上？
这里我们就要用到优先级队列（二叉堆）这种数据结构，把链表放入一个最小堆，就可以每次获得K个结点的最小节点。
*/

//最小堆
type minHeap []*ListNode

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//解题
func mergeNLists(lists []*ListNode) *ListNode {
	h := new(minHeap)
	if len(lists) > 1 {
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				heap.Push(h, lists[i])
			}
		}
		dummyHead := new(ListNode)
		p := dummyHead
		for h.Len() > 0 {
			tmp := heap.Pop(h).(*ListNode)
			if tmp.Next != nil {
				heap.Push(h, tmp.Next)
			}
			p.Next = tmp
			p = p.Next
		}
		return dummyHead.Next

	} else if len(lists) == 1 {
		return lists[0]
	} else {
		return nil
	}
}
