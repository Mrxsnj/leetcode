/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var in func(root *TreeNode)
	in = func(root *TreeNode) {
		if root != nil {
			in(root.Left)
			res = append(res, root.Val)
			in(root.Right)
		}
	}

	in(root)
	return res
}