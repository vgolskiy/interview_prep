package src

func MajorElement(nums []int) int {
	count := 0
	current := 0
	for _, n := range nums {
		if count == 0 {
			current = n
		}

		if n == current {
			count++
		} else {
			count--
		}
	}
	return current
}
