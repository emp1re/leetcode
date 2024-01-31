package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	curr := head
	runner := head

	for curr.Next != nil && runner.Next != nil {
		runner = runner.Next.Next
		if runner != nil {
			curr = curr.Next
			if runner == curr {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
