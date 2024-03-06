package main

import "testing"

func ExpectedOutput(head *ListNode, expected []int) (bool, []int) {
	var actual []int
	retVal := true

	for i, tmp := 0, head; tmp != nil; i, tmp = i+1, tmp.Next {
		if tmp.Val != expected[i] {
			retVal = false
		}

		actual = append(actual, tmp.Val)
	}

	return retVal, actual
}

func TestRev00(t *testing.T) {
	var head *ListNode

	head = reverseKGroup(head, 2)

	expected := []int{}
	if ok, actual := ExpectedOutput(head, expected); !ok {
		t.Error("Expected: ", expected, " Got: ", actual)
	}
}

func TestRev0(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	head = reverseKGroup(head, 2)

	expected := []int{1}
	if ok, actual := ExpectedOutput(head, expected); !ok {
		t.Error("Expected: ", expected, " Got: ", actual)
	}
}

func TestRev1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	head = reverseKGroup(head, 2)

	expected := []int{2, 1, 4, 3, 5}
	if ok, actual := ExpectedOutput(head, expected); !ok {
		t.Error("Expected: ", expected, " Got: ", actual)
	}
}

func TestRev2(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	head = reverseKGroup(head, 1)

	expected := []int{1, 2, 3, 4, 5}
	if ok, actual := ExpectedOutput(head, expected); !ok {
		t.Error("Expected: ", expected, " Got: ", actual)
	}
}

func TestRev3(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	head = reverseKGroup(head, 2)

	expected := []int{2, 1, 4, 3, 6, 5}
	if ok, actual := ExpectedOutput(head, expected); !ok {
		t.Error("Expected: ", expected, " Got: ", actual)
	}
}

func TestRev4(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}

	head = reverseKGroup(head, 3)

	expected := []int{3, 2, 1, 6, 5, 4}
	if ok, actual := ExpectedOutput(head, expected); !ok {
		t.Error("Expected: ", expected, " Got: ", actual)
	}
}
