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
	st := []*TreeNode{}

	for {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}

		if len(st) == 0 {
			break
		} else {
			root = st[len(st)-1].Right
			res = append(res, st[len(st)-1].Val)
			st = st[:len(st)-1]
		}
	}

	return res
}