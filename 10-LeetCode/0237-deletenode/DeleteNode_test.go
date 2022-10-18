package deletenode

import (
	"example/gophercises/v2/10-LeetCode/ln"
	"testing"
)

func deleteNode(node *ln.ListNode) {
	if node.Next != nil {
		// copy the value from the next node to this node
		node.Val = node.Next.Val

		// have this node skip the next node and go on to
		// next's next
		node.Next = node.Next.Next
	}
}

// ----- tests ----- //

func Test_DeleteNode_S001(t *testing.T) {
	testDeleteNode(t, "4519", 5, "419")
}

func Test_DeleteNode_S002(t *testing.T) {
	testDeleteNode(t, "4519", 1, "459")
}

// ----- helpers ----- //

func testDeleteNode(t *testing.T, list string, node int, expected string) {

	// arrange: create the list
	var head *ln.ListNode = ln.StringToList(list)

	// arrange: find the node we want to delete
	var n = head
	for n != nil {
		if n.Val == node {
			break
		} else {
			n = n.Next
		}
	}

	if n == nil {
		t.Errorf("Could not find node '%d'!", node)
	}

	// act: attempt to delete the node
	deleteNode(n)

	// assert: confirm the deletion was successful
	var actual string = ln.ListToString(head)
	if expected != actual {
		t.Errorf("Expected=%s Actual=%s", expected, actual)
	}
}
