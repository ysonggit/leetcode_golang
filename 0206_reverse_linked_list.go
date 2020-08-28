func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		nex := cur.Next // pre  cur->nex->
		cur.Next = pre  // pre<-cur  nex->
		pre = cur       // pre==cur  nex->
		cur = nex       //      pre  cur->
	}
	return pre
}