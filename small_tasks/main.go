package main

import (
	"InterviewPrep/src"
	"fmt"
)

func main() {
	fmt.Println(src.Hello())

	fmt.Println(src.Reverse("!ti esreveR"))

	arr := [5]int{1, 2, 3, 4, 5}
	sli := arr[1:3]
	fmt.Println(sli)
	sli = append(sli, 6)
	fmt.Println(sli)
	fmt.Println(arr)
	sli = append(sli, 7)
	fmt.Println(sli)
	fmt.Println(arr)
	sli = append(sli, []int{8, 9}...)
	fmt.Println(sli)
	fmt.Println(arr)

	lst := []int{1, 2, 3, 4}
	fmt.Println(src.Two_sum(lst, 6))

	fmt.Println(1001, src.IsPalindrome(1001))
	fmt.Println(-1001, src.IsPalindrome(-1001))
	fmt.Println(123, src.IsPalindrome(123))
	fmt.Println(0, src.IsPalindrome(0))

	fmt.Println("XXIV - 24", src.RomanToInt("XXIV"))
	fmt.Println("MCMXLIV - 1944", src.RomanToInt("MCMXLIV"))

	fmt.Println("FizzBizz:", src.FizzBizz(sli))

	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	fmt.Println("Merge two non-descending (might be similar values) arrays:", src.Merge(nums1, m, nums2, n))

	n, num := src.RemoveElement([]int{2, 2, 3}, 2)
	fmt.Println("Remove element:", n, num)

	n, num = src.RemoveDuplicates([]int{1, 2, 3})
	fmt.Println("Remove duplicates: [1,2,3]", n, num)

	fmt.Println("Majority element in the array: [2,2,1,1,1,2,2]", src.MajorElement([]int{2, 2, 1, 1, 1, 2, 2}))

	fmt.Println("Read file by chunks:")
	src.ReadFileByChunks()

	fmt.Println("The highest profit: [7,1,5,3,6,4]", src.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
}
