package ln

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func StringToList(str string) *ListNode {
	var head *ListNode = new(ListNode)
	var current = head

	for index, r := range str {
		var i, err = strconv.Atoi(string(r))
		if err != nil {
			panic(i)
		}
		current.Val = i
		if index < len(str)-1 {
			current.Next = new(ListNode)
			current = current.Next
		}
	}

	return head
}

func ListToString(head *ListNode) string {
	var str string = ""
	for head != nil {
		str = fmt.Sprintf("%s%d", str, head.Val)
		head = head.Next
	}
	return str
}
