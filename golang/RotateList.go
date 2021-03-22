package main

import (
	"fmt"
	"strconv"
)

func toList(s []int) *ListNode {
	var p, t *ListNode
	for _, v := range s {
		if p == nil {
			p = &ListNode{v, nil}
			t = p
		} else {
			t.Next = &ListNode{v, nil}
			t = t.Next
		}
	}
	return p
}

func (p *ListNode) String() string {
	if p.Next == nil {
		return strconv.Itoa(p.Val)
	} else {
		return strconv.Itoa(p.Val) + "," + p.Next.String()
	}
}

func length(head *ListNode) int {
	if head == nil {
		return 0
	}
	sum := 1
	for v := head; v.Next != nil; v = v.Next {
		sum++
	}
	return sum
}

func rotateRight(head *ListNode, k int) *ListNode {
	l := length(head)
	if l == 0 {
		return nil
	}
	k = k % l
	p := head
	for i := 0; i < l-k-1; i++ {
		p = p.Next
	}
	newHead := p.Next
	if newHead == nil {
		return head
	} else {
		p.Next = nil
		t := newHead
		for t.Next != nil {
			t = t.Next
		}
		t.Next = head
		return newHead
	}
}

func main() {
	//fmt.Println(rotateRight(toList([]int{1, 2, 3, 4, 5}), 1))
	//fmt.Println(rotateRight(toList([]int{1, 2, 3, 4, 5}), 2))
	//fmt.Println(rotateRight(toList([]int{0, 1, 2}), 1))
	//fmt.Println(rotateRight(toList([]int{0, 1, 2}), 2))
	//fmt.Println(rotateRight(toList([]int{0, 1, 2}), 3))
	//fmt.Println(rotateRight(toList([]int{0, 1, 2}), 4))
	fmt.Println(rotateRight(toList([]int{}), 1))
}
