package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	var tmp []int
	inter_map := make(map[int]bool)

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for _, val := range nums1 {
		inter_map[val] = false
	}

	for _, val := range nums2 {
		if _, ok := inter_map[val]; ok {
			inter_map[val] = true
		}
	}

	for k, v := range inter_map {
		if v {
			tmp = append(tmp, k)
		}
	}

	return tmp
}

func main() {
	s1 := []int{1, 2, 2, 1}
	s2 := []int{2, 2}

	out := intersection(s1, s2)

	fmt.Println(out)
}
