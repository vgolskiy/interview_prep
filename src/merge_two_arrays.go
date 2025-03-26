package src

// going from the end, as we will be able to save numbers on the right position right away
func save(nums1 *[]int, m *int, nums2 *[]int, n *int, pos *int) {
	if (*nums1)[*m] > (*nums2)[*n] {
		(*nums1)[*pos] = (*nums1)[*m]
		*m--
	} else if (*nums1)[*m] < (*nums2)[*n] {
		(*nums1)[*pos] = (*nums2)[*n]
		*n--
	} else {
		(*nums1)[*pos] = (*nums1)[*m]
		*m--
		*pos--
		(*nums1)[*pos] = (*nums2)[*n]
		*n--
	}
	*pos--
}

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	pos := m + n - 1
	m--
	n--

	for m >= 0 && n >= 0 {
		save(&nums1, &m, &nums2, &n, &pos)
	}

	for n >= 0 {
		nums1[n] = nums2[n]
		n--
	}
	return nums1
}
