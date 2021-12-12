package chunk

func ProcessLine(line string) (bool, rune) {
	opens := make([]rune, 0)
	for _, r := range line[0:] {
		switch r {
		case '(', '[', '{', '<':
			opens = append(opens, r)
		case ')', ']', '}', '>':
			open := opens[len(opens) - 1]
			expectedClose := openClose[open]
			if r != expectedClose{
				return false, r
			}
			opens = opens[:len(opens) - 1]
		}
	}
	return true, 0
}

func Complete(line string) []rune {
	opens := make([]rune, 0)
	for _, r := range line[0:] {
		switch r {
		case '(', '[', '{', '<':
			opens = append(opens, r)
		case ')', ']', '}', '>':
			open := opens[len(opens) - 1]
			expectedClose := openClose[open]
			if r != expectedClose{
				panic("tried to complete invalid line")
			}
			opens = opens[:len(opens) - 1]
		}
	}
	closes := make([]rune, len(opens))
	for i := 0; i < len(opens); i++ {
		openValue := opens[len(opens) - 1 - i]
		closes[i] = openClose[openValue]
	}
	return closes
}

func ScoreCloses(closes []rune) int {
	total := 0
	for _, close := range closes {
		total *= 5
		total += CloseValues[close]
	}
	return total
}

var openClose = map[rune]rune {
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var Values = map[rune]int {
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var CloseValues = map[rune]int {
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}
