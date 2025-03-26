package src

var converter = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var substruct = map[rune]rune{
	'V': 'I',
	'X': 'I',
	'L': 'X',
	'C': 'X',
	'D': 'C',
	'M': 'C',
}

func RomanToInt(s string) int {
	runes := []rune(s)

	result := 0
	for i := 0; i < len(runes); {
		if i+1 < len(runes) {
			combo := substruct[runes[i+1]]
			if combo == runes[i] {
				result += converter[runes[i+1]] - converter[runes[i]]
				i += 2
				continue
			}
		}
		result += converter[runes[i]]
		i++
	}
	return result
}
