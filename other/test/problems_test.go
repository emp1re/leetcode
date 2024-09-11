package test

import (
	"testing"

	"github.com/emp1re/leetcode/other/internal/problems"
	"github.com/stretchr/testify/assert"
)

func TestProblems(t *testing.T) {
	t.Logf("%s: %+v", "2894. Divisible and Non-divisible Sums Difference.go", problems.DifferenceOfSums(10, 3))
	t.Logf("%s: %s", "1108. Defanging an IP Address.go", problems.DefangIPaddr("192.168.0.1"))
	t.Logf("%s: %+v", "1637. Widest Vertical Area Between Two Points Containing No Points.go", problems.MaxWidthOfVerticalArea([][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}))
	t.Logf("%s: %+v", "1470. Shuffle the Array.go", problems.Shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	t.Logf("%s: %+v", "1920. Build Array from Permutation.go", problems.BuildArray([]int{0, 2, 1, 5, 3, 4}))
	t.Logf("%s: %+v", "1929. Concatenation of Array.go", problems.GetConcatenation([]int{1, 2, 3}))
	t.Logf("%s: %+v", "2413. Smallest Even Multiple.go", problems.SmallestEvenMultiple(3))
	t.Logf("%s: %f", "2469. Convert the Temperature.go", problems.ConvertTemperature(32))
	t.Logf("%s: %+v", "2769. Find the Maximum Achievable Number.go", problems.TheMaximumAchievableX(20, 3))
	t.Logf("%s: %+v", "2798. Number of Employees Who Met the Target.go", problems.NumberOfEmployeesWhoMetTarget([]int{1, 2, 3, 4, 5}, 3))
	t.Logf("%s: %+v", "2942. Find Words Containing Character.go", problems.FindWordsContaining([]string{"hello", "world", "leetcode"}, 'e'))

	t.Logf("%s: %+v", "1365. How Many Numbers Are Smaller Than the Current Number.go", problems.SmallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))

	t.Logf("%s: %+v", "1512. Number of Good Pairs.go", problems.NumIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	t.Logf("%s: %+v", "1672. Richest Customer Wealth.go", problems.MaximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	t.Logf("%s: %+v", "2974. Minimum Number Game.go", problems.NumberGame([]int{1, 2, 3, 4, 5, 6}))
	t.Logf("%s: %+v", "1480. Running Sum of 1d Array.go", problems.RunningSum([]int{1, 2, 3, 4}))
	t.Logf("%s: %+v", "2114. Maximum Number of Words Found in Sentences.go", problems.MostWordsFound([]string{"hello world", "hello world", "hello world", "hello world"}))
	t.Logf("%s: %+v", "1282.", problems.GroupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	t.Logf("%s: %+v", "1313. Decompress Run-Length Encoded List.go", problems.DecompressRLElist([]int{1, 2, 3, 4}))
	t.Logf("%s: %+v", "118. Pascal's Triangle.go", problems.Generate(5))
	t.Logf("%s: %+v", "53. Maximum Subarray.go", problems.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	t.Logf("%s: %+v", "1389. Create Target Array in the Given Order.go", problems.CreateTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	t.Logf("%s: %+v", "1221. Split a String in Balanced Strings.go", problems.BalancedStringSplit("RLRRLLRLRL"))
	t.Logf("%s: %s", "2810. Faulty Keyboard.go", problems.FinalString("a#b#C#d#e#f#G#h#i#j#k#L#m#n#O#p#q#r#S#t#U#v#W#x#Y#z#"))
	t.Logf("%s: %+v", "2660. Determine the Winner of a Bowling Game.go", problems.IsWinner([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 5}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}))
	t.Logf("%s: %+v", "771. Jewels and Stones.go", problems.NumJewelsInStones("aA", "aAAbbbb"))
	t.Logf("%s: %+v", "2006. Count Number of Pairs With Absolute Difference.go", problems.CountKDifference([]int{1, 2, 2, 1}, 1))
	t.Logf("%s: %+v", "804. Unique Morse Code Words.go", problems.UniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
	t.Logf("%s: %+v", "2520. Count the Digits That Divide a Number", problems.CountDigits(1012))
	t.Logf("%s: %+v", "1342. Number of Steps to Reduce a Number to Zero.go", problems.NumberOfSteps(14))
	t.Logf("%s: %+v", "1486. XOR Operation in an Array.go", problems.XorOperation(5, 0))
	t.Logf("%s: %+v", "2433. Find The Original Array of Prefix Xor", problems.FindArray([]int{6, 2, 7, 3}))
	t.Logf("%s: %+v", "2367. Number of Arithmetic Triplets.go", problems.ArithmeticTriplets([]int{1, 2, 4, 6, 7, 10}, 3))
	t.Logf("%s: %+v", "238. Missing Number.go", problems.MissingNumber([]int{3, 0, 1}))
	t.Logf("%s: %+v", "283. Move Zeroes.go", assert.Equal(t, []int{1, 3, 12, 0, 0}, problems.MoveZeroes([]int{0, 1, 0, 3, 12})))
	t.Logf("%s: %+v", "1502. Can Make Arithmetic Progression From Sequence", assert.Equal(t, true, problems.CanMakeArithmeticProgression([]int{3, 5, 1})))
}
