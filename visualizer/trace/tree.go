package trace

import "fmt"

// TreeNode is a binary tree node used by traced solutions.
type TreeNode struct {
	ID    string
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree parses a LeetCode-format array into a linked tree and a flat node list.
// Input: []any where elements are int or nil.
// Returns the root TreeNode and a slice of Node for rendering.
func BuildTree(input []any) (*TreeNode, []Node) {
	if len(input) == 0 {
		return nil, nil
	}

	// Parse first element
	rootVal, ok := toInt(input[0])
	if !ok {
		return nil, nil
	}

	root := &TreeNode{ID: "n0", Val: rootVal}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(input) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(input) {
			if v, ok := toInt(input[i]); ok {
				leftID := fmt.Sprintf("n%d", i)
				node.Left = &TreeNode{ID: leftID, Val: v}
				queue = append(queue, node.Left)
			}
			i++
		}

		// Right child
		if i < len(input) {
			if v, ok := toInt(input[i]); ok {
				rightID := fmt.Sprintf("n%d", i)
				node.Right = &TreeNode{ID: rightID, Val: v}
				queue = append(queue, node.Right)
			}
			i++
		}
	}

	// Compute layout positions and build flat node list
	nodes := layoutTree(root)
	return root, nodes
}

// layoutTree assigns x,y positions using inorder traversal for x and level for y.
func layoutTree(root *TreeNode) []Node {
	var nodes []Node
	xCounter := 0.0
	const ySpacing = 80.0

	var inorder func(node *TreeNode, depth int)
	inorder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		inorder(node.Left, depth+1)

		n := Node{
			ID:  node.ID,
			Val: node.Val,
			X:   xCounter * 100,
			Y:   float64(depth) * ySpacing,
		}
		if node.Left != nil {
			n.LeftID = node.Left.ID
		}
		if node.Right != nil {
			n.RightID = node.Right.ID
		}
		nodes = append(nodes, n)
		xCounter++

		inorder(node.Right, depth+1)
	}
	inorder(root, 0)
	return nodes
}

// toInt converts an any value to int, returns false for nil.
func toInt(v any) (int, bool) {
	switch val := v.(type) {
	case int:
		return val, true
	case float64:
		return int(val), true
	case nil:
		return 0, false
	default:
		return 0, false
	}
}
