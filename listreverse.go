package piscine

func ListReverse(l *List) {
	var prev *NodeL = nil
	curr := l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head, l.Tail = prev, l.Head
}
