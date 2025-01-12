package main

import (
	"fmt"
	"math"
)

func main() {
	node1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	node2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}

	res := addTwoNumbers(node1, node2)
	fmt.Println(res)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current1 := l1
	current2 := l2
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	for current1 != nil {
		arr1 = append(arr1, current1.Val)
		current1 = current1.Next
	}
	for current2 != nil {
		arr2 = append(arr2, current2.Val)
		current2 = current2.Next
	}
	mergedArr := make([]int, 0)
	if len(arr1) >= len(arr2) {
		forNext := 0
		for i := 0; i < len(arr1); i++ {
			currentNum := arr1[i] + forNext
			if len(arr2) >= i+1 {
				currentNum += arr2[i]
			}
			if currentNum < 10 {
				mergedArr = append(mergedArr, currentNum)
				forNext = 0
			} else {
				num := float64(currentNum) / 10
				forNext = int(num)
				forAdd := int(math.Round((num - float64(forNext)) * 10))
				mergedArr = append(mergedArr, forAdd)
			}
		}
		if forNext > 0 {
			mergedArr = append(mergedArr, forNext)
		}
	} else if len(arr2) >= len(arr1) {
		forNext := 0
		for i := 0; i < len(arr2); i++ {
			currentNum := arr2[i] + forNext
			if len(arr1) >= i+1 {
				currentNum += arr1[i]
			}
			if currentNum < 10 {
				mergedArr = append(mergedArr, currentNum)
				forNext = 0
			} else {
				num := float64(currentNum) / 10
				forNext = int(num)
				forAdd := int(math.Round((num - float64(forNext)) * 10))
				mergedArr = append(mergedArr, forAdd)
			}
		}
		if forNext > 0 {
			mergedArr = append(mergedArr, forNext)
		}
	}
	fmt.Println(mergedArr)
	result := &ListNode{Val: mergedArr[0], Next: nil}
	current := result
	for i := 0; i < len(mergedArr); i++ {
		if i > 0 {
			node := &ListNode{Val: mergedArr[i], Next: nil}
			current.Next = node
			current = node
		}
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}
