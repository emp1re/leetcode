package test

import (
	"testing"

	"github.com/emp1re/leetcode/other/internal/problems"
)

func TestProblems(t *testing.T) {
	t.Logf("%s: %d", "2894. Divisible and Non-divisible Sums Difference.go", problems.DifferenceOfSums(10, 3))
	t.Logf("%s: %s", "1108. Defanging an IP Address.go", problems.DefangIPaddr("192.168.0.1"))
	t.Logf("%s: %d", "1637. Widest Vertical Area Between Two Points Containing No Points.go", problems.MaxWidthOfVerticalArea([][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}))
	t.Logf("%s: %d", "1470. Shuffle the Array.go", problems.Shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	t.Logf("%s: %d", "1920. Build Array from Permutation.go", problems.BuildArray([]int{0, 2, 1, 5, 3, 4}))
	t.Logf("%s: %d", "1929. Concatenation of Array.go", problems.GetConcatenation([]int{1, 2, 3}))
	t.Logf("%s: %d", "2413. Smallest Even Multiple.go", problems.SmallestEvenMultiple(3))
	t.Logf("%s: %f", "2469. Convert the Temperature.go", problems.ConvertTemperature(32))
	t.Logf("%s: %d", "2769. Find the Maximum Achievable Number.go", problems.TheMaximumAchievableX(20, 3))
	t.Logf("%s: %d", "2798. Number of Employees Who Met the Target.go", problems.NumberOfEmployeesWhoMetTarget([]int{1, 2, 3, 4, 5}, 3))
	t.Logf("%s: %d", "2942. Find Words Containing Character.go", problems.FindWordsContaining([]string{"hello", "world", "leetcode"}, 'e'))
	t.Logf("%s: %d", "1365. How Many Numbers Are Smaller Than the Current Number.go", problems.SmallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}
