package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	new := l.Head
	for new != nil {
		if comp(new.Data, ref) {
			return &new.Data
		}
		new = new.Next
	}
	return nil
}
