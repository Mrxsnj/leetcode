/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var pre func(root *TreeNode)
	pre = func(root *TreeNode) {
		if root != nil {
			res = append(res, root.Val)
			pre(root.Left)
			pre(root.Right)
		}
	}

	pre(root)
	return res
}

