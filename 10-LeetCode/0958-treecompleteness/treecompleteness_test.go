package treecompleteness

import (
	"example/gophercises/v2/10-LeetCode/ln"
	"testing"
)

func isCompleteTree(root *ln.TreeNode) bool {
	numComplete := countCompleteTree(root, 0, 0)
	return numComplete <= 1
}

func countCompleteTree(root *ln.TreeNode, curSum int, depth int) int {
	if root.Left == nil && root.Right == nil {
		// leafs sum to zero
		return curSum + 0
	} else if root.Left != nil && root.Right == nil {
		// simple case - also a complete tree
		return curSum + 1
	} else if root.Left == nil && root.Right != nil {
		// simple case - not a complete tree
		return curSum + 2
	} else {
		// have two branches, count each of the branches
		leftSum := countCompleteTree(root.Left, curSum, depth+1)

		// if we already found a complete tree to the left AND
		// have more to the right, then it's not a complete tree
		// and can just shortcut with a large number here
		if leftSum >= 1 {
			return 2
		}
		return countCompleteTree(root.Right, leftSum, depth+1)
	}
}

func TestExampleOne(t *testing.T) {
	runScenario("1,2,3,4,5,6", true, t)
}

func TestExampleTwo(t *testing.T) {
	runScenario("1,2,3,4,5,null,7", false, t)
}

func Test73(t *testing.T) {
	runScenario("1,2,3,5,null,7,8", false, t)
}

func Test73a(t *testing.T) {
	runScenario("1,2,3,5,null,7,null", false, t)
}

func Test92(t *testing.T) {
	runScenario("1,2", true, t)
}

func Test104(t *testing.T) {
	runScenario("1,2,3,5", true, t)
}

func Test105(t *testing.T) {
	runScenario("1,2,3,null,null,7,8", false, t)
}

func runScenario(tree string, expected bool, t *testing.T) {
	t.Helper()

	root := ln.BuildTree(tree)

	var actual = isCompleteTree(root)

	if expected != actual {
		t.Errorf("Expected: %v Actual: %v", expected, actual)
	}
}
