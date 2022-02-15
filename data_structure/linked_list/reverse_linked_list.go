package linked_list

/**
  @Author: pika
  @Date: 2022/2/11 11:12 上午
  @Description: 反转链表：「反转整个链表」、「反转链表前n个节点」、「反转链表」
*/

/*
 对于递归算法，最重要的就是明确递归函数的定义。具体来说，我们的reverse函数定义是这样的：
 输入一个节点head，将「以head为起点」的链表反转，并返回反转后的头结点
*/

/*
 1、递归函数要有 base case，也就是这句：
	if (head.next == null) return head;
 意思是如果链表只有一个节点的时候反转也是它自己，直接返回即可。

 2、当链表递归反转之后，新的头结点是 last，而之前的 head 变成了最后一个节点，别忘了链表的末尾要指向 null：
	head.next = null;

*/

//反转链表节点
func reverseLinkedList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := reverseLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//反转链表前N个节点
var successor *ListNode //后驱节点

func reverseLinkedListN(head *ListNode, n int) *ListNode {
	if n == 1 {
		//记录n+1个节点
		successor = head.Next
		return head
	}
	// 以 head.Next为起点，需要反转前 n-1个节点
	last := reverseLinkedListN(successor, n-1)
	head.Next.Next = head
	//让反转之后的head节点和后面的节点连起来
	head.Next = successor
	return last
}

//反转链表的一部分
//给一个索引区间【m,n】,仅仅反转区间中的链表元素

func reverseLinkedListBetween(head *ListNode, b, e int) *ListNode {
	//base case
	if b == 1 {
		//相当于反转前n个链表
		return reverseLinkedListN(head, e)
	}

	//前进到反转的起点触发 base case
	head.Next = reverseLinkedListBetween(head, b-1, e-1)

	return head
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/*
	k个一组反转链表
	这个问题具有递归性质
*/

//反转a节点至b节点之间的内容[a,b)左闭右开
func reverseBetAB(a *ListNode, b *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	pre = nil
	cur, nxt = a, a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	//返回反转后的头结点
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	//区间[start, end)包含k个待反转元素
	var start, end *ListNode
	start, end = head, head
	for i := 0; i < k; i++ {
		if end == nil {
			return start
		}
		end = end.Next
	}
	//反转前k个元素，不包含end，左闭右开，所以end未被反转
	newHead := reverseBetAB(start, end)
	// 递归反转后续链表并连接起来
	start.Next = reverseKGroup(end, k)

	return newHead
}
