package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	list := SplitDaList(str)
	listfull := make(map[string]int)
	for _, item := range list {
		listfull[item] += 1
	}
	return listfull
}

func SplitDaList(s string) []string {
	var sentence []string
	word := ""
	for _, letter := range s {
		if letter == ' ' {
			sentence = append(sentence, word)
			word = ""
		} else {
			word += string(letter)
		}
	}
	sentence = append(sentence, word)
	return sentence
}
