package main

import (
	"os"
)

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

func TrimAtoi(s string) int {
	integ := 0
	flag := false
	for _, i := range s {
		if i == '-' && integ == 0 {
			flag = true
		}
		if i >= '0' && i <= '9' {
			integ *= 10
			integ += int(i - '0')
		}
	}
	if flag {
		integ *= -1
	}
	return integ
}

func main() {
	args := os.Args
	if len(args) < 1 {
		return
	}
	ta := os.Args[1]
	tb := os.Args[2]
	tc := os.Args[3]
	c := 0
	i := 0
	for i < len(ta) {
		if c == 0 && ta[i] == '-' {
			i += 1
		}
		if ta[i] < '0' || ta[i] > '9' {
			return
		}
		i += 1
	}
	c = 0
	i = 0
	for i < len(tc) {
		if c == 0 && tc[i] == '-' {
			i += 1
		}
		if tc[i] < '0' || tc[i] > '9' {
			return
		}
		i += 1
	}
	taint := TrimAtoi(ta)
	tcint := TrimAtoi(tc)
	if tb == "+" || tb == "-" || tb == "/" || tb == "%" || tb == "*" {
		var result int
		switch tb {
		case "+":
			result = taint + tcint
		case "-":
			result = taint - tcint
		case "*":
			result = taint * tcint
		case "/":
			if tcint == 0 {
				newstring := "No division by 0 \n"
				os.Stdout.Write([]byte(newstring))
				return
			}
			result = taint / tcint
		case "%":
			if tcint == 0 {
				newstring := "No modulo by 0 \n"
				os.Stdout.Write([]byte(newstring))
				return
			}
			result = taint % tcint
		}
		newstring := Atoi(result)
		newstring += "\n"
		if result > -9223372036854775808 && result < 9223372036854775807 {
			os.Stdout.Write([]byte(newstring))
		}
	} else {
		return
	}
}
