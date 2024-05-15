package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}
	for {
		unsorted := false
		var prev *NodeI = nil
		curr := l
		next := l.Next
		for next != nil {
			if curr.Data > next.Data {
				if prev != nil {
					prev.Next = next
				} else {
					l = next
				}
				unsorted = true
				curr.Next = next.Next
				next.Next = curr
				prev = next
				next = curr.Next
			} else {
				prev = curr
				curr = next
				next = next.Next
			}
		}
		if !unsorted {
			break
		}
	}
	return l
}
