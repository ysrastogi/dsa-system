package maximumwidhttree
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
type NodeIndex struct {
    node  *TreeNode
    index int
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func widthOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

    q := []NodeIndex{{root, 0}}
    maxWidth := 0

    for len(q) > 0 {
        size := len(q)
        start := q[0].index
        var end int

        for i := 0; i < size; i++ {
            front := q[0]
            q = q[1:]

            node := front.node
            idx := front.index - start // normalize
            end = idx

            if node.Left != nil {
                q = append(q, NodeIndex{node.Left, 2*idx + 1})
            }
            if node.Right != nil {
                q = append(q, NodeIndex{node.Right, 2*idx + 2})
            }
        }

        maxWidth = max(maxWidth, end+1)
    }

    return maxWidth
}
