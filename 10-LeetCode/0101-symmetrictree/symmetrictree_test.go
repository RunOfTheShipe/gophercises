package symmettrictree

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return nodesEqual(root.Left, root.Right)
}

func nodesEqual(left *TreeNode, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right == nil {
		return true
	} else {
		if left.Val == right.Val {
			return nodesEqual(left.Left, right.Right) && nodesEqual(left.Right, right.Left)
		}
		return false
	}
}

// TESTS
func TestExampleOne(t *testing.T) {
	input := []int{1, 2, 2, 3, 4, 4, 3}
	testScenario(input, true, t)
}

func testScenario(tree []int, expected bool, t *testing.T) {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
