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
	st := []*TreeNode{}
	var lastPrinted *TreeNode

	for {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}

		if len(st) == 0 {
			break
		} else {
			root = st[len(st)-1]
		}

		if root.Right == nil || root.Right == lastPrinted {
			lastPrinted = root
			res = append(res, root.Val)
			st = st[:len(st)-1]
			root = nil
		} else {
			root = root.Right
		}
	}

	return res
}