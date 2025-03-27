package src

func Hello() string {
	return "Hello World"
}

func Reverse(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}
