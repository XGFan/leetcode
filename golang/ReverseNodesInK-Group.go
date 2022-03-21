package main

import (
	"fmt"
	"strconv"
)

//Suppress all
type ListNode struct {
	Val  int
	Next *ListNode
}

func (p *ListNode) String() string {
	if p.Next == nil {
		return strconv.Itoa(p.Val)
	} else {
		return strconv.Itoa(p.Val) + "," + p.Next.String()
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	ans, tail := reverseFirstK(head, k)
	for p := tail.Next; p != nil; {
		tail.Next, tail = reverseFirstK(p, k)
		p = tail.Next
	}
	return ans
}
func reverseFirstK(head *ListNode, k int) (*ListNode, *ListNode) {
	var i = 1
	p := head.Next
	var newHead, tail = head, head
	tail.Next = nil
	for ; i < k && p != nil; i++ {
		newHead, p, p.Next = p, p.Next, newHead
	}
	if p != nil {
		tail.Next = p
	}
	if i < k {
		return reverseFirstK(newHead, i)
	}
	return newHead, tail
}

//func reverseKGroup(head *ListNode, k int) *ListNode {
//	p := head
//	var ans *ListNode = nil
//	var t *ListNode = nil
//	for p != nil {
//		newHead := reverseOrNot(p, k) //newHead 为头，p为尾巴
//		if ans == nil {               //第一次遍历
//			ans = newHead //ans 指向头
//		} else {
//			t.Next = newHead //连上
//		}
//		t = p //t指向尾巴
//		p = p.Next
//	}
//	return ans
//}
func reverseOrNot(head *ListNode, size int) *ListNode {
	h := head
	var p *ListNode
	if head.Next != nil {
		p = head.Next
	}
	var i = 1
	for p != nil && i < size {
		head.Next = p.Next //head实际上是tail
		p.Next = h         //当前点p的next指向旧head h
		h = p              //新的head h指向p
		p = head.Next      //p指向下一个点
		i++
	}
	if i < size {
		return reverseOrNot(h, i) //如果长度不够，干脆再翻转一次
	} else {
		return h
	}

}

/**
fun reverseKGroup(head: ListNode?, k: Int): ListNode? {
    var p = head
    var ans: ListNode? = null
    var t: ListNode? = null
    while (p != null) {
        val newHead = reverseOrNot(p, k) //newHead 为头，p为尾巴
        if (ans == null) { //第一次遍历
            ans = newHead //ans 指向头
        } else {
            t?.next = newHead //连上
        }
        t = p //t指向尾巴
        p = p.next
    }
    return ans
}

fun reverseOrNot(head: ListNode?, size: Int): ListNode? {
    var h = head
    var p = head?.next
    var i = 1
    while (p != null && i < size) {
        head?.next = p.next //head实际上是tail
        p.next = h //当前点p的next指向旧head h
        h = p //新的head h指向p
        p = head?.next //p指向下一个点
        i++
    }
    return if (i < size) {
        reverseOrNot(h, i) //如果长度不够，干脆再翻转一次
    } else {
        h
    }
}
*/

func main() {
	list := &ListNode{Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		}}
	//k, l := reverseFirstK(list, 2)
	//fmt.Printf("[%s]\t[%s]\n", k, l)
	k1 := reverseKGroup(list, 2)
	fmt.Printf("[%s]\n", k1)
}
