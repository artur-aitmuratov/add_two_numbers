package main

import "fmt"

func main() {
	l1 := ListNode[1, _],
	l2 = ListNode{1, *l1}
	addTwoNumbers(&l1, &l2)
}

func addTwoNumbers(nums1 *ListNode, nums2 *ListNode) {
	fmt.Println(nums1, nums2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
