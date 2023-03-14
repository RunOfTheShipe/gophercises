/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package sumroottoleaf

import (
	"fmt"
	"strconv"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sumRecurse(root, "", 0)
}

func sumRecurse(current *TreeNode, txt string, sum int) int {
	if current == nil {
		return 0
	} else if current.Left == nil && current.Right == nil {
		txt = fmt.Sprintf("%s%d", txt, current.Val)

		// convert txt to number (specify base 10 to cover the
		// case of leading zeros)
		num, _ := strconv.ParseInt(txt, 10, 32)

		// return the number
		return int(sum) + int(num)
	} else {
		txt = fmt.Sprintf("%s%d", txt, current.Val)
		return sumRecurse(current.Left, txt, sum) + sumRecurse(current.Right, txt, sum)
	}
}

func Test100(t *testing.T) {

	root := new(TreeNode)
	root.Val = 0

	r := new(TreeNode)
	r.Val = 0
	root.Right = r

	temp := new(TreeNode)
	temp.Val = 6
	r.Right = temp
	r = temp

	temp = new(TreeNode)
	temp.Val = 8
	r.Right = temp
	r = temp

	temp = new(TreeNode)
	temp.Val = 7
	r.Right = temp
	r = temp

	sum := sumNumbers(root)
	if sum != 687 {
		t.Errorf("Expected 687, Actual %d", sum)
	}

}
