// https://leetcode.com/problems/binary-tree-preorder-traversal/
package leetcode

import (
	. "lc/utils"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorder(node *TreeNode, result []int) []int {
	if node != nil {
		result = AppendUnique(result, []int{node.Val})
		result = AppendUnique(result, preorder(node.Left, result))
		result = AppendUnique(result, preorder(node.Right, result))
		return result
	} else {
		return []int{}
	}
}

func preorderTraversal(root *TreeNode) []int {
	return preorder(root, []int{})
}

// Input
// root =
// [37,-34,-48,null,-100,-100,48,null,null,null,null,-54,null,-71,-22,null,null,null,8]

// Use Testcase
// Output
// [37,-34,-100,-48,48,-54,-71,-22,8]
// Expected
// [37,-34,-100,-48,-100,48,-54,-71,-22,8]
