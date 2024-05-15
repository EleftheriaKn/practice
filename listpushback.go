package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	new := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = new
		l.Tail = new
	} else {
		l.Tail.Next = new
		l.Tail = new
	}
}
