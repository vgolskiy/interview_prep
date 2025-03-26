package src

import (
	"strconv"
)

func FizzBizz(arr []int) []string {
	var res []string

	for _, n := range arr {
		if n%3 == 0 && n%5 == 0 {
			res = append(res, "FizzBizz")
		} else if n%3 == 0 {
			res = append(res, "Fizz")
		} else if n%5 == 0 {
			res = append(res, "Bizz")
		} else {
			res = append(res, strconv.Itoa(n))
		}
	}
	return res
}
