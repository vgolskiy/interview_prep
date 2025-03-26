package src

func IsPalindrome(a int) bool {
	if a >= 0 {
		original := a
		reverse := 0

		for a > 0 {
			reverse = reverse*10 + a%10
			a /= 10
		}
		return original == reverse
	}
	return false
}
