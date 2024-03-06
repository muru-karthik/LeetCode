package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(start, end *ListNode) (*ListNode, *ListNode) {
	var prev, curr, next *ListNode
	tmp := start

	if start != nil {
		prev, curr, next = start, start, start

		if prev.Next == nil { // only one node
			return prev, prev
		}
		curr = prev.Next
		prev.Next = nil
		next = curr.Next
	} else { // Empty list
		return nil, nil
	}

	for {
		curr.Next = prev
		if curr == end {
			return curr, tmp
		} else {
			prev = curr
			curr = next
			next = curr.Next
		}
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	tmp := head
	start, end := head, head
	prev := head

	firstIter := true

	if k == 1 {
		return head
	}

	for i := 0; ; tmp, i = tmp.Next, i+1 {
		if i == k {
			next := tmp
			start, end = reverse(start, end)
			if firstIter {
				firstIter = false
				head = start
			}
			prev.Next = start
			end.Next = next
			prev = end
			start, tmp = next, next
			i = 0
		}
		end = tmp

		if tmp == nil {
			break
		}
	}

	return head
}

func main() {
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
}
