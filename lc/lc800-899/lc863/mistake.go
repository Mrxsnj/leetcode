/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
    flag int
    depth int
    res []int
)

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    flag = 0
    depth = 0
    res = []int{}

    dfs(root, target, k)

    return res
}

func dfs(root *TreeNode, target *TreeNode, k int) {
    if root == nil {
        return
    }

    if root.Val == target.Val {
        flag = 1
    } else {
        // else for: don't count target itself
        if flag == 1 {
            depth++
        } else if flag == -1 {
            depth--
        }
    }
    
    if depth == k || depth == -k {
        res = append(res, root.Val)
    }

    dfs(root.Left, target, k)

    if depth == k || depth == -k {
        res = append(res, root.Val)
    }
    
    dfs(root.Right, target, k)

    if depth == k || depth == -k {
        res = append(res, root.Val)
    }

    if root.Val == target.Val {
        flag = -1
        depth--
    } else {
        if flag == 1 {
            depth--
        } else if flag == -1 {
            depth++
        }
    }
}
