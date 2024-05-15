package piscine

func Compact(ptr *[]string) int {
	length := len(*ptr) - 1
	for i := 0; i < length; {
		if (*ptr)[i] == "" {
			copy(((*ptr)[i:]), (*ptr)[i+1:])
			length--
		} else {
			i++
		}
	}
	(*ptr) = (*ptr)[:length]
	return length
}
