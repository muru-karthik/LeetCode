package main

import (
	"sort"
	"testing"
)

func compare(i, j int) bool {
	return i < j
}

func checkOutput(n1, n2 []int) bool {
	if len(n1) != len(n2) {
		return false
	}
	sort.Slice(n1, compare)
	sort.Slice(n2, compare)

	for idx, val := range n1 {
		if val != n2[idx] {
			return false
		}
	}

	return true
}

func TestCase1(t *testing.T) {
	s1 := []int{1, 2, 2, 1}
	s2 := []int{2, 2}

	expected := []int{2}
	out := intersection(s1, s2)

	if !checkOutput(expected, out) {
		t.Errorf("Failed. Expected:%v, Actual:%v\n", expected, out)
	}
}
