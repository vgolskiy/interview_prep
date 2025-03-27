package src

func LongestCommonPrefix(strs []string) string {
	base := []rune(strs[0])
	for _, w := range strs[1:] {
		if len(base) == 0 || len(w) == 0 {
			return ""
		}
		runes := []rune(w)
		maxLength := min(len(runes), len(base))
		base = base[:maxLength]
		for i, r := range runes[:maxLength] {
			if base[i] != r {
				base = runes[:i]
				break
			}
		}
	}
	return string(base)
}
