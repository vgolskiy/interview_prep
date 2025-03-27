package src

func Two_sum(nums []int, prod int) []int {
	mapper := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mapper[nums[i]]; ok {
			return []int{mapper[nums[i]], i}
		} else {
			mapper[prod-nums[i]] = i
		}
	}
	return []int{}
}
