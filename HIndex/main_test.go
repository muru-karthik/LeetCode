package main

import "testing"

func TestHindex1(t *testing.T) {
	citations := []int{1, 0, 5, 3, 6}

	hindex := hIndex(citations)

	if hindex != 3 {
		t.Fatalf("Wanted:3 , Got:%v", hindex)
	}
}

func TestHindex2(t *testing.T) {
	citations := []int{1, 3, 1}

	hindex := hIndex(citations)

	if hindex != 1 {
		t.Fatalf("Wanted:1 , Got:%v", hindex)
	}
}

func TestHindex3(t *testing.T) {
	citations := []int{100}

	hindex := hIndex(citations)

	if hindex != 1 {
		t.Fatalf("Wanted:1 , Got:%v", hindex)
	}
}

func TestHindex4(t *testing.T) {
	citations := []int{11, 15}

	hindex := hIndex(citations)

	if hindex != 2 {
		t.Fatalf("Wanted:2 , Got:%v", hindex)
	}
}
