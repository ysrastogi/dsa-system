package childernsumproperty

import (
	"fmt"

	"github.com/yashrastogi1/dsa-system/visualizer/trace"
)

// ChildrenSumTraced runs the children sum property algorithm with full execution tracing.
func ChildrenSumTraced(input []any) trace.Trace {
	tracer := trace.NewTracer("children-sum-property", "push-down-pull-up")
	root, nodes := trace.BuildTree(input)
	tracer.SetTree(input, nodes)

	if root == nil {
		return tracer.Export(nil)
	}

	// Track node states for highlights
	nodeStates := make(map[string]string)
	for _, n := range nodes {
		nodeStates[n.ID] = "unvisited"
	}

	// Track current values for display
	nodeVals := make(map[string]int)
	var initVals func(n *trace.TreeNode)
	initVals = func(n *trace.TreeNode) {
		if n == nil {
			return
		}
		nodeVals[n.ID] = n.Val
		initVals(n.Left)
		initVals(n.Right)
	}
	initVals(root)

	buildHighlights := func(currentID string, currentState string) []trace.Highlight {
		old := nodeStates[currentID]
		nodeStates[currentID] = currentState
		var hl []trace.Highlight
		for _, n := range nodes {
			hl = append(hl, trace.Highlight{NodeID: n.ID, State: nodeStates[n.ID]})
		}
		if currentState == "current" {
			nodeStates[currentID] = old
		}
		return hl
	}

	// Emit initial state
	tracer.Emit("init", root.ID,
		map[string]any{"phase": "start"},
		buildHighlights(root.ID, "unvisited"),
		"Tree loaded — will apply Children Sum Property",
	)

	var modify func(node *trace.TreeNode)
	modify = func(node *trace.TreeNode) {
		if node == nil {
			return
		}

		tracer.PushCall(fmt.Sprintf("modify(%d)", node.Val))

		// Visit
		tracer.Emit("visit", node.ID,
			map[string]any{"node": node.Val, "phase": "compare"},
			buildHighlights(node.ID, "current"),
			fmt.Sprintf("Enter modify(%d)", node.Val),
		)
		nodeStates[node.ID] = "instack"

		// Phase 1: Compare
		childSum := 0
		if node.Left != nil {
			childSum += node.Left.Val
		}
		if node.Right != nil {
			childSum += node.Right.Val
		}

		if childSum >= node.Val {
			// Children are already big enough
			oldVal := node.Val
			node.Val = childSum
			nodeVals[node.ID] = node.Val
			tracer.Emit("compute", node.ID,
				map[string]any{
					"node":     node.Val,
					"childSum": childSum,
					"phase":    "compare",
				},
				buildHighlights(node.ID, "instack"),
				fmt.Sprintf("childSum=%d ≥ node=%d → set node to %d", childSum, oldVal, childSum),
			)
		} else {
			// Push parent's value down to children
			tracer.Emit("update", node.ID,
				map[string]any{
					"node":     node.Val,
					"childSum": childSum,
					"phase":    "push-down",
				},
				buildHighlights(node.ID, "instack"),
				fmt.Sprintf("childSum=%d < node=%d → PUSH DOWN %d to children", childSum, node.Val, node.Val),
			)
			if node.Left != nil {
				node.Left.Val = node.Val
				nodeVals[node.Left.ID] = node.Val
			}
			if node.Right != nil {
				node.Right.Val = node.Val
				nodeVals[node.Right.ID] = node.Val
			}
		}

		// Phase 2: Recurse
		modify(node.Left)
		modify(node.Right)

		// Phase 3: Pull up
		total := 0
		if node.Left != nil {
			total += node.Left.Val
		}
		if node.Right != nil {
			total += node.Right.Val
		}
		if node.Left != nil || node.Right != nil {
			oldVal := node.Val
			node.Val = total
			nodeVals[node.ID] = total
			nodeStates[node.ID] = "computed"
			tracer.Emit("return", node.ID,
				map[string]any{
					"node":  total,
					"phase": "pull-up",
				},
				buildHighlights(node.ID, "computed"),
				fmt.Sprintf("Pull up: %d → %d (sum of children's final values)", oldVal, total),
			)
		} else {
			nodeStates[node.ID] = "done"
			tracer.Emit("return", node.ID,
				map[string]any{
					"node":  node.Val,
					"phase": "leaf",
				},
				buildHighlights(node.ID, "done"),
				fmt.Sprintf("Leaf node — value stays %d", node.Val),
			)
		}

		tracer.PopCall()
	}

	modify(root)

	// Collect final values
	finalVals := make(map[string]int)
	var collectFinal func(n *trace.TreeNode)
	collectFinal = func(n *trace.TreeNode) {
		if n == nil {
			return
		}
		finalVals[n.ID] = n.Val
		collectFinal(n.Left)
		collectFinal(n.Right)
	}
	collectFinal(root)

	return tracer.Export(finalVals)
}
