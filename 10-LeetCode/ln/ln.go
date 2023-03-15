package ln

import (
	"fmt"
	"strconv"
	"strings"
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(tree string) *TreeNode {
	vals := strings.Split(tree, ",")
	treeData := make([]*int, len(vals))

	for i, val := range vals {
		if val == "null" || val == "nil" {
			treeData[i] = nil
		} else {
			data, _ := strconv.Atoi(val)
			treeData[i] = &data
		}
	}

	return buildTree(&treeData)
}

func buildTree(tree *[]*int) *TreeNode {
	root := new(TreeNode)
	root.Val = *(*tree)[0]

	build(root, 0, tree)

	return root
}

func build(root *TreeNode, index int, treeData *[]*int) {
	leftIndex := 2*index + 1
	rightIndex := 2*index + 2

	if leftIndex < len(*treeData) && (*treeData)[leftIndex] != nil {
		root.Left = new(TreeNode)
		root.Left.Val = *(*treeData)[leftIndex]

		build(root.Left, leftIndex, treeData)
	}

	if rightIndex < len(*treeData) && (*treeData)[rightIndex] != nil {
		root.Right = new(TreeNode)
		root.Right.Val = *(*treeData)[rightIndex]

		build(root.Right, rightIndex, treeData)
	}
}

func Int(val int) *int {
	return &val
}
