package test

import (
	"testing"

	"github.com/emp1re/leetcode/top150/internal/problems"
)

func TestProblems(t *testing.T) {
	list := problems.ListNode{Val: 1, Next: &problems.ListNode{Val: 2, Next: &problems.ListNode{Val: 3, Next: &problems.ListNode{Val: 4, Next: &problems.ListNode{Val: 5, Next: nil}}}}}
	t.Logf("%s: %t", "141. Linked List Cycle.go", problems.HasCycle(&list))
	t.Logf("%s: %d", "21. Merge Two Sorted Lists.go", problems.MergeTwoLists(&problems.ListNode{Val: 1, Next: &problems.ListNode{Val: 2, Next: &problems.ListNode{Val: 4, Next: nil}}}, &problems.ListNode{Val: 1, Next: &problems.ListNode{Val: 3, Next: &problems.ListNode{Val: 4, Next: nil}}}))

}
func BenchmarkMinimumSizeSubarraySum(b *testing.B) {
	bigSet := []int{2, 3, 1, 2, 4, 3}
	b.SetBytes(2)
	test := problems.MinSubArrayLen(396893380, bigSet)

	b.Log("result :", test)
}
