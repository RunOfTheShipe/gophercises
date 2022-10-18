package addtwo

import (
	"fmt"
	"strconv"
	"testing"
)

func TestExampleOne(t *testing.T) {
	testAddTwoNumber(t, "243", "564", "708")
}

func TestExampleTwo(t *testing.T) {
	testAddTwoNumber(t, "0", "0", "0")
}

func TestExampleThree(t *testing.T) {
	testAddTwoNumber(t, "9999999", "9999", "89990001")
	// l1 		9999999
	// l2 		9999
	// Expected=89990001
	//   Actual=89991010101
}

func TestExampleFour(t *testing.T) {
	testAddTwoNumber(t, "1000000000000000000000000000001", "564", "6640000000000000000000000000001")
}

func testAddTwoNumber(t *testing.T, l1 string, l2 string, expected string) {
	var list1 *ListNode = stringToList(l1)
	var list2 *ListNode = stringToList(l2)

	var answer *ListNode = addTwoNumbers(list1, list2)

	var actual string = listToString(answer)

	if expected != actual {
		t.Errorf("Expected=%s Actual=%s", expected, actual)
	}
}

func stringToList(str string) *ListNode {
	var head = new(ListNode)
	var current = head

	for index, r := range str {
		var v, err = strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		current.Val = v

		if index != len(str)-1 {
			current.Next = new(ListNode)
			current = current.Next
		}
	}

	return head
}

func listToString(l *ListNode) string {
	var str string = ""
	for l != nil {
		str = fmt.Sprintf("%s%d", str, l.Val)
		l = l.Next
	}
	return str
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var answer *ListNode = new(ListNode)

	addLists(l1, l2, answer, nil, false)

	return answer
}

func addLists(l1 *ListNode, l2 *ListNode, answer *ListNode, previous *ListNode, carry bool) {

	if l1 == nil && l2 == nil {
		// end case
		if carry {
			answer.Val = 1
			answer.Next = nil
		} else if previous != nil {
			// nothing got carried over and we don't want a
			// leading zero, so need to null out the previous
			// pointer to the current answer
			previous.Next = nil
		}
	} else if l1 != nil && l2 != nil {

		var sum int = l1.Val + l2.Val
		if carry {
			sum += 1
		}
		answer.Val = sum % 10
		answer.Next = new(ListNode)
		addLists(l1.Next, l2.Next, answer.Next, answer, sum > 9)
	} else if l1 != nil { // l2 is empty
		var sum int = l1.Val
		if carry {
			sum += 1
		}

		answer.Val = sum % 10
		answer.Next = new(ListNode)
		addLists(l1.Next, nil, answer.Next, answer, sum > 9)
	} else { // l1 is empty, l2 has contents
		var sum int = l2.Val
		if carry {
			sum += 1
		}

		answer.Val = sum % 10
		answer.Next = new(ListNode)
		addLists(nil, l2.Next, answer.Next, answer, sum > 9)
	}
}
