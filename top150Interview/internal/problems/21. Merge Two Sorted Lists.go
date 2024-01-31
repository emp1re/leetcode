package problems

import "fmt"

func len(list *ListNode) int {
	count := 0
	cur := list
	for cur != nil {
		cur = cur.Next
		count++
	}
	return count
}
func append(list *ListNode, val int) bool {
	newNode := &ListNode{Val: val, Next: nil}
	cur := list
	for cur != nil {
		newNode.Next = cur.Next
		cur.Next = newNode

		return true
	}
	return false
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curr := list1
	curr2 := list2
	if list1 == nil {
		return list2
	}
	for curr2 != nil {
		append(list1, curr2.Val)
		curr2 = curr2.Next
	}
	for i := 0; i < len(list1); i++ {
		curr = list1
		for curr.Next != nil {
			if curr.Val > curr.Next.Val {

				curr.Val, curr.Next.Val = curr.Next.Val, curr.Val
				fmt.Println(curr.Val, curr.Next.Val)
			}
			curr = curr.Next
		}
	}
	return list1
}
