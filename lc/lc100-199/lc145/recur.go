/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var post func(root *TreeNode)
	post = func(root *TreeNode) {
		if root != nil {
			post(root.Left)
			post(root.Right)
			res = append(res, root.Val)
		}
	}

	post(root)
	return res
}