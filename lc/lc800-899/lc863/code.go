/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
    res []int
    parents map[int]*TreeNode
)

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    res = []int{}
    parents = map[int]*TreeNode{}

    findParents(root)
    dfs(target, nil, k, 0)

    return res
}

func findParents(root *TreeNode) {
    if root == nil {
        return
    }

    if root.Left != nil {
        parents[root.Left.Val] = root
        findParents(root.Left)
    }

    if root.Right != nil {
        parents[root.Right.Val] = root
        findParents(root.Right)
    }
}

func dfs(target *TreeNode, from *TreeNode, k, depth int) {
    if target == nil {
        return
    }
    
    if depth == k {
        res = append(res, target.Val)
    }

    if target.Left != from {
        dfs(target.Left, target, k, depth+1)
    }

    if target.Right != from {
        dfs(target.Right, target, k, depth+1)
    }

    if parents[target.Val] != from {
        dfs(parents[target.Val], target, k, depth+1)
    }
}
