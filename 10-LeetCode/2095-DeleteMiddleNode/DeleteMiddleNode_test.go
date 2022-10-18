package deletemiddlenode

import (
	"example/gophercises/v2/10-LeetCode/ln"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ln.ListNode) *ln.ListNode {

	var moveMiddlePtr bool = false

	var middle *ln.ListNode = head
	var previousMiddle *ln.ListNode = nil
	var current *ln.ListNode = head

	for current != nil {
		// always move current to the next one
		current = current.Next

		// only move the middle on every other move
		if moveMiddlePtr {
			previousMiddle = middle
			middle = middle.Next
		}

		// always flip this
		moveMiddlePtr = !moveMiddlePtr
	}

	// previous middle is pointing to the one just before
	// the middle; use that to skip over the middle
	if previousMiddle == nil {
		// list of one item
		head = nil
	} else {
		previousMiddle.Next = middle.Next
	}

	return head
}

func deleteMiddle_Slow(head *ln.ListNode) *ln.ListNode {

	// count the number of nodes in the list
	var current = head
	var count int = 0
	for current != nil {
		count++
		current = current.Next
	}

	// determine the middle (divide by two)
	var middle int = count / 2

	// go back to middle, skip it
	current = head
	var previous *ln.ListNode = nil
	for i := 0; i < middle; i++ {
		previous = current
		current = current.Next
	}

	// skip the current node node
	previous.Next = current.Next

	return head
}

// ----- tests ----- //
func Test_DeleteMiddle_S001(t *testing.T) {
	testDeleteMiddle(t, "1347126", "134126")
}

func Test_DeleteMiddle_S002(t *testing.T) {
	testDeleteMiddle(t, "1234", "124")
}

func Test_DeleteMiddle_S003(t *testing.T) {
	testDeleteMiddle(t, "21", "2")
}

func Test_DeleteMiddle_S004(t *testing.T) {
	testDeleteMiddle(t, "1", "")
}

// ----- helpers ----- //

func testDeleteMiddle(t *testing.T, list string, expected string) {
	var input *ln.ListNode = ln.StringToList(list)

	var output *ln.ListNode = deleteMiddle(input)

	var actual string = ln.ListToString(output)
	if expected != actual {
		t.Errorf("Expected=%s Actual=%s", expected, actual)
	}
}
