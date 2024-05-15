package piscine

func Atoi(num int) string {
	if num == 0 {
		return "0"
	}

	negative := false
	if num < 0 {
		negative = true
		num = -num
	}

	length := 0
	tmp := num
	for tmp > 0 {
		length++
		tmp /= 10
	}

	result := make([]byte, length)
	for i := length - 1; i >= 0; i-- {
		result[i] = byte('0' + num%10)
		num /= 10
	}

	if negative {
		return "-" + string(result)
	}
	return string(result)
}
