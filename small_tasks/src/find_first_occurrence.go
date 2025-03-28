package src

func FirstOccurrence(haystack string, needle string) int {
	word := []rune(needle)
	base := []rune(haystack)
	j := 0
	for i := 0; i < len(base); i++ {
		for j < len(word) && i < len(base) {
			if base[i] == word[j] {
				i++
				j++
			} else {
				i = i - j
				j = 0
				break
			}
		}
		if j == len(word) {
			return i - len(word)
		}
	}
	return -1
}
