func insert(aNode *Node, x int) *Node {
	node := &Node{
		Val: x,
	}
	if aNode == nil {
		node.Next = node
		return node
	}
	for cur := aNode.Next; true; cur = cur.Next {
		next := cur.Next
		if aNode == cur || (cur.Val <= x && next.Val >= x) || (next.Val < cur.Val && (x > cur.Val || x < next.Val)) {
			node.Next = next
			cur.Next = node
			break
		}
	}
	return aNode
}