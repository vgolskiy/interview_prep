package src

func RemoveDuplicates(nums []int) (int, []int) {
	n := 0
	pos := 0
	catch := true
	last_pos := len(nums) - 1

	for i := 0; i < last_pos; {
		for i < last_pos && nums[i] == nums[i+1] {
			if catch {
				pos = i
				catch = false
			}
			i++
		}
		for i < last_pos && nums[i] != nums[i+1] {
			if pos != i {
				nums[pos] = nums[i]
			}
			pos++
			i++
			n++
		}
	}

	if pos != last_pos {
		nums[pos] = nums[last_pos]
	}
	n++
	return n, nums
}
