package maxpathsum

import (
	"fmt"

	"github.com/yashrastogi1/dsa-system/visualizer/trace"
)

// MaxPathSumTraced runs the max path sum algorithm with full execution tracing.
func MaxPathSumTraced(input []any) trace.Trace {
	tracer := trace.NewTracer("max-path-sum", "dfs-postorder")
	root, nodes := trace.BuildTree(input)
	tracer.SetTree(input, nodes)

	if root == nil {
		return tracer.Export(0)
	}

	globalMax := root.Val

	// Build highlights helper — marks all nodes with their current state
	nodeStates := make(map[string]string) // nodeID → state
	for _, n := range nodes {
		nodeStates[n.ID] = "unvisited"
	}

	buildHighlights := func(currentID string, currentState string) []trace.Highlight {
		old := nodeStates[currentID]
		nodeStates[currentID] = currentState
		var hl []trace.Highlight
		for _, n := range nodes {
			hl = append(hl, trace.Highlight{NodeID: n.ID, State: nodeStates[n.ID]})
		}
		if currentState == "current" {
			nodeStates[currentID] = old // restore for non-permanent states
		}
		return hl
	}

	// Emit initialization step
	tracer.Emit("init", root.ID,
		map[string]any{"globalMax": globalMax},
		buildHighlights(root.ID, "unvisited"),
		fmt.Sprintf("Initialize globalMax = %d (root value)", globalMax),
	)

	var dfs func(node *trace.TreeNode) int
	dfs = func(node *trace.TreeNode) int {
		if node == nil {
			return 0
		}

		tracer.PushCall(fmt.Sprintf("dfs(%d)", node.Val))

		// Visit step
		tracer.Emit("visit", node.ID,
			map[string]any{"globalMax": globalMax, "node": node.Val},
			buildHighlights(node.ID, "current"),
			fmt.Sprintf("Enter dfs(%d)", node.Val),
		)
		nodeStates[node.ID] = "instack"

		// Recurse left
		left := 0
		if node.Left != nil {
			left = max(0, dfs(node.Left))
		}

		// Recurse right
		right := 0
		if node.Right != nil {
			right = max(0, dfs(node.Right))
		}

		// Compute forked path
		forked := left + right + node.Val
		returnVal := max(left, right) + node.Val

		// Check for global max update
		if forked > globalMax {
			globalMax = forked
			nodeStates[node.ID] = "computed"
			tracer.Emit("update", node.ID,
				map[string]any{
					"globalMax": globalMax,
					"left":      left,
					"right":     right,
					"forked":    forked,
					"return":    returnVal,
				},
				buildHighlights(node.ID, "computed"),
				fmt.Sprintf("🏆 NEW GLOBAL MAX! forked = %d + %d + %d = %d", left, right, node.Val, forked),
			)
		} else {
			nodeStates[node.ID] = "computed"
			tracer.Emit("compute", node.ID,
				map[string]any{
					"globalMax": globalMax,
					"left":      left,
					"right":     right,
					"forked":    forked,
					"return":    returnVal,
				},
				buildHighlights(node.ID, "computed"),
				fmt.Sprintf("forked = %d + %d + %d = %d (no update, global=%d)", left, right, node.Val, forked, globalMax),
			)
		}

		// Return step
		nodeStates[node.ID] = "done"
		tracer.Emit("return", node.ID,
			map[string]any{
				"globalMax": globalMax,
				"return":    returnVal,
			},
			buildHighlights(node.ID, "done"),
			fmt.Sprintf("Return %d to parent (best single chain upward)", returnVal),
		)

		tracer.PopCall()
		return returnVal
	}

	dfs(root)
	return tracer.Export(globalMax)
}
