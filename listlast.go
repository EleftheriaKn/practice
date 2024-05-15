package piscine

func ListLast(l *List) interface{} {
	last := l.Tail
	if last == nil {
		return nil
	}
	return last.Data
}
