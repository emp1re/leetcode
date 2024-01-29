package main

import (
	"fmt"
	"log/slog"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	slog.Info("Start...", "main.go")
	l := []int{1, 2, 3, 4, 5}
	var result *ListNode
	for _, v := range l {
		if result == nil {
			result = &ListNode{Val: v}
		} else {
			result.add(v)
		}
	}
	fmt.Println(result)
}
func (l *ListNode) add(value int) {
	newNode := &ListNode{Val: value}
	if l.Next == nil {
		l.Next = newNode
	} else {
		l.Next.add(value)
	}

}
