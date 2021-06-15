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

	st := []*TreeNode{}

	for {
		for root != nil {
			res = append(res, root.Val)
			st = append(st, root)
			root = root.Left
		}

		if len(st) == 0 {
			break
		} else {
			root = st[len(st)-1].Right
			st = st[:len(st)-1]
		}
	}

	// for len(st) > 0 || root != nil {
	// 	for root != nil {
	// 		res = append(res, root.Val)
	// 		st = append(st, root)
	// 		root = root.Left
	// 	}

	// 	root = st[len(st)-1].Right
	// 	st = st[:len(st)-1]
	// }

	return res
}