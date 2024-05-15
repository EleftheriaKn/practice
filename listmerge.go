package piscine

func ListMerge(l1 *List, l2 *List) {
	// new := l1.Head
	// for new != nil {
	// 	new = new.Next
	// }
	// newl2 := l2.Head
	// new = newl2
	// for newl2 != nil {
	// 	l1.Tail = newl2
	// }
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	} else if l2.Head != nil {
		l1.Tail.Next = l2.Head
		l1.Tail = l2.Tail
	}
}
