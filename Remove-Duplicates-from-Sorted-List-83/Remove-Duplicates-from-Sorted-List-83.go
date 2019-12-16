package main

import "fmt"

func main() {

	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: &ListNode{},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(deleteDuplicates(&head))
}

// //Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	preNode := head
	currentNode := head.Next
	for currentNode != nil {
		if preNode.Val == currentNode.Val {
			preNode.Next = currentNode.Next
		} else {
			preNode = currentNode
		}
		currentNode = currentNode.Next
	}

	return head
}
