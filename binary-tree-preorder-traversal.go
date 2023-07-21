// https://leetcode.com/problems/binary-tree-preorder-traversal/
package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
	Sum []int
}

func NewTree(root *TreeNode) Tree {
	return Tree{ Root: root }
}

func (tree *Tree) preorder(node *TreeNode) {
	if node != nil {
		tree.Sum = append(tree.Sum, node.Val)
		tree.preorder(node.Left)
		tree.preorder(node.Right)
	}
}

func (tree *Tree) Preorder() {
	tree.preorder(tree.Root)
}

func preorderTraversal(root *TreeNode) []int {
	tree := NewTree(root)
	tree.Preorder()
	return tree.Sum
}

