package listtest

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList1(count int) *ListNode {
	dumb := &ListNode{}
	p := dumb
	cur := 0
	for cur < count {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = cur + 1
		cur += 1
	}
	return dumb.Next
}

func CreateList2(values []int) *ListNode {
	if values == nil {
		return nil
	}
	dummy := &ListNode{}
	p := dummy
	for _, val := range values {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = val
	}
	return dummy.Next
}

func printList(list *ListNode) {
	if list == nil {
		return
	}
	p := list
	for p != nil {
		fmt.Printf(" %d", p.Val)
		p = p.Next
	}
	fmt.Println("")
}

func FindLastK(list *ListNode, K int) *ListNode {
	if list == nil || K <= 0 {
		return nil
	}
	p := list
	q := list
	count := 0
	for count < K && p != nil {
		p = p.Next
		count++
	}
	if count < K {
		return nil
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q
}

func TestFindLastK(t *testing.T) {
	list1 := CreateList1(6)
	printList(list1)
	p := FindLastK(list1, 2)
	if p == nil {
		t.Log("p == nil")
	} else {
		t.Logf("p.data: %d", p.Val)
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	pre := dummy
	cur := head

	for cur != nil {
		next := cur.Next
		if next == nil {
			pre.Next = cur
			break
		}
		for next != nil && cur.Val == next.Val {
			next = next.Next
		}

		if next == nil {
			pre.Next = nil
			break
		}
		if cur.Next == next {
			pre.Next = cur
			pre = cur
		}
		cur = next
	}
	return dummy.Next
}

func TestDeleteDuplicates(t *testing.T) {
	// s1 := []int{1, 2, 3, 3, 4, 4, 5}
	s1 := []int{1, 2, 2}
	list1 := CreateList2(s1)
	printList(list1)
	list1 = deleteDuplicates(list1)
	t.Log("after dup")
	printList(list1)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy

	p := l1
	q := l2
	carry := 0
	for p != nil && q != nil {
		node := &ListNode{}
		node.Val = (p.Val + q.Val + carry) % 10
		carry = (p.Val + q.Val + carry) / 10
		pre.Next = node
		pre = node
		p = p.Next
		q = q.Next
	}

	for p != nil {
		node := &ListNode{}
		node.Val = (p.Val + carry) % 10
		carry = (p.Val + carry) / 10
		pre.Next = node
		pre = node
		p = p.Next
	}

	for q != nil {
		node := &ListNode{}
		node.Val = (q.Val + carry) % 10
		carry = (q.Val + carry) / 10
		pre.Next = node
		pre = node
		q = q.Next
	}

	if carry > 0 {
		node := &ListNode{}
		node.Val = carry
		pre.Next = node
		pre = node
	}

	return dummy.Next
}
