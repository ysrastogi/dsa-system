package templates

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TreeDP performs postorder DFS, computing a value bottom-up.
// - Each node receives results from children, combines them, returns to parent.
// - A global variable tracks the best "through this node" answer.
//
// Pattern: What flows UP (return), what is GLOBAL (side-effect update).
//
// Usage: Max path sum, diameter, longest univalue path.

var globalMax int

// MaxPathSum returns the maximum path sum in a binary tree.
// A path can start and end at any node.
func MaxPathSum(root *TreeNode) int {
	globalMax = root.Val
	dfsMaxPath(root)
	return globalMax
}

func dfsMaxPath(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// Only take positive contributions
	left := max(0, dfsMaxPath(node.Left))
	right := max(0, dfsMaxPath(node.Right))

	// Global: path that forks through this node
	globalMax = max(globalMax, left+right+node.Val)

	// Return: best single chain upward (can't fork)
	return max(left, right) + node.Val
}
