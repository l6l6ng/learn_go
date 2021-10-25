package main

import "fmt"

func main() {
	//fmt.Println(solve([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	p4 := ListNode{Val: 4, Next: nil}
	p3 := ListNode{Val: 3, Next: &p4}
	p2 := ListNode{Val: 2, Next: &p3}
	p := ListNode{Next: &p2}
	pHead := &p

	fmt.Println(ReverseList(pHead))
}

func solve(a []int) int {
	// write code here
	len := len(a)
	if len == 1 {
		return 0
	}

	for i := len - 1; i > 0; i-- {
		if i == len-1 && a[i] >= a[i-1] {
			return i
		} else if i < len-1 && (a[i] >= a[i-1] && a[i] >= a[i+1]) {
			return i
		}
	}

	return 0
}

func judge(str string) bool {
	// write code here
	len := len(str)
	if len > 1 {
		for i := 0; i < len/2; i++ {
			if str[i] != str[len-i-1] {
				return false
			}
		}
	}
	return true
}

func solve2(a []int) int {
	// write code here
	k := 0
	for i, v := range a {
		k ^= (i + 1) ^ v
	}
	return k
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead.Next != nil {
		var pNew *ListNode

		p := pHead.Next
		for p != nil {
			t := p.Next
			p.Next = pNew.Next
			pNew.Next = p
			p = t
		}
		return pNew
	}
	return nil
}
