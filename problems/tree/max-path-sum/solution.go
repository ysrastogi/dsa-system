package maxpathsum

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxPathSum returns the maximum path sum in a binary tree.
//
// Key insight: two different values at each node.
//   - GLOBAL: the best "forked" path through this node (left + right + node)
//   - RETURN: the best "chain" going upward (max(left, right) + node)
//
// Time: O(n) — visit every node once
// Space: O(h) — recursion stack depth (h = height)
func MaxPathSum(root *TreeNode) int {
	globalMax := root.Val
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// Recurse into children — clamp negatives to 0
		left := max(0, dfs(node.Left))
		right := max(0, dfs(node.Right))

		// GLOBAL update: path forking through this node
		// This is left-chain + node + right-chain
		forked := left + right + node.Val
		if forked > globalMax {
			globalMax = forked
		}

		// RETURN to parent: best single chain upward
		// Parent can only use ONE of our subtrees (path can't fork twice)
		return max(left, right) + node.Val
	}

	dfs(root)
	return globalMax
}
