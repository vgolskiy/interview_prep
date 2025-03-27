package src

func RemoveElement(nums []int, val int) (int, []int) {
	n := 0
	pos := 0
	catch := true
	for i := 0; i < len(nums); {
		for i < len(nums) && nums[i] == val {
			if catch {
				pos = i
				catch = false
			}
			i++
		}
		for i < len(nums) && nums[i] != val {
			nums[pos] = nums[i]
			pos++
			i++
			n++
		}
	}
	return n, nums
}
