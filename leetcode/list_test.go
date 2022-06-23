package leetcode

import (
	"fmt"
	"testing"
)

type ListNode struct {
	data int
	next *ListNode
}

func createList(count int) *ListNode {
	dumb := &ListNode{}
	p := dumb
	cur := 0
	for cur < count {
		p.next = &ListNode{}
		p = p.next
		p.data = cur + 1
		cur += 1
	}
	return dumb.next
}

func printList(list *ListNode) {
	if list == nil {
		return
	}
	p := list
	for p != nil {
		fmt.Printf(" %d", p.data)
		p = p.next
	}
	fmt.Println("")
}

func findLastK(list *ListNode, K int) *ListNode {
	if list == nil || K <= 0 {
		return nil
	}
	p := list
	q := list
	count := 0
	for count < K && p != nil {
		p = p.next
		count++
	}
	if count < K {
		return nil
	}
	for p != nil {
		p = p.next
		q = q.next
	}
	return q
}

func TestList1(t *testing.T) {
	list1 := createList(6)
	printList(list1)
	p := findLastK(list1, 2)
	if p == nil {
		t.Log("p == nil")
	} else {
		t.Logf("p.data: %d", p.data)
	}
}
