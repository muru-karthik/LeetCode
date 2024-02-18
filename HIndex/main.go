package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	hindex := 0
	sort.Ints(citations)
	totalPubs := len(citations)

	if totalPubs == 1 && citations[0] >= 1 {
		return 1
	}
	for idx, val := range citations {
		if val >= totalPubs-idx {
			hindex = min(val, idx+1)
		}
	}
	return hindex
}

func main() {
	citations := []int{1, 3, 0, 5, 6}
	citations = []int{1, 3, 1}
	hindex := hIndex(citations)

	fmt.Printf("H-Index:%v\n", hindex)
}
