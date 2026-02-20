package trace

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Step represents one atomic action in the algorithm execution.
type Step struct {
	ID         int            `json:"id"`
	Action     string         `json:"action"`     // "visit", "return", "compare", "update", "init"
	NodeID     string         `json:"nodeId"`     // which node this step is about
	Variables  map[string]any `json:"variables"`  // snapshot of all tracked variables
	Highlights []Highlight    `json:"highlights"` // which nodes to highlight and how
	Message    string         `json:"message"`    // human-readable explanation
	CallStack  []string       `json:"callStack"`  // current call stack
}

// Highlight specifies how to visually mark a node.
type Highlight struct {
	NodeID string `json:"nodeId"`
	State  string `json:"state"` // "current", "instack", "computed", "done", "pathnode", "unvisited"
}

// Node is the renderable tree node with layout positions.
type Node struct {
	ID      string  `json:"id"`
	Val     int     `json:"val"`
	LeftID  string  `json:"leftId,omitempty"`
	RightID string  `json:"rightId,omitempty"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
}

// Trace is the full execution trace for a problem.
type Trace struct {
	Problem   string `json:"problem"`
	Algorithm string `json:"algorithm"` // "dfs-postorder", "bfs-level", "sliding-window"
	TreeInput []any  `json:"treeInput"` // LeetCode format input
	Nodes     []Node `json:"nodes"`     // tree structure for renderer
	Steps     []Step `json:"steps"`
	Answer    any    `json:"answer"`
}

// Tracer records execution steps.
type Tracer struct {
	problem   string
	algorithm string
	treeInput []any
	nodes     []Node
	steps     []Step
	callStack []string
	stepID    int
}

// NewTracer creates a new tracer for a problem.
func NewTracer(problem, algorithm string) *Tracer {
	return &Tracer{
		problem:   problem,
		algorithm: algorithm,
	}
}

// SetTree stores the tree input and computed nodes.
func (t *Tracer) SetTree(input []any, nodes []Node) {
	t.treeInput = input
	t.nodes = nodes
}

// PushCall adds a function call to the call stack.
func (t *Tracer) PushCall(name string) {
	t.callStack = append(t.callStack, name)
}

// PopCall removes the top call from the call stack.
func (t *Tracer) PopCall() {
	if len(t.callStack) > 0 {
		t.callStack = t.callStack[:len(t.callStack)-1]
	}
}

// Emit records a step in the trace.
func (t *Tracer) Emit(action, nodeID string, vars map[string]any, highlights []Highlight, msg string) {
	t.stepID++
	// Copy the call stack
	stack := make([]string, len(t.callStack))
	copy(stack, t.callStack)

	// Copy the variables map
	varsCopy := make(map[string]any, len(vars))
	for k, v := range vars {
		varsCopy[k] = v
	}

	t.steps = append(t.steps, Step{
		ID:         t.stepID,
		Action:     action,
		NodeID:     nodeID,
		Variables:  varsCopy,
		Highlights: highlights,
		Message:    msg,
		CallStack:  stack,
	})
}

// Export finalizes the trace with the answer.
func (t *Tracer) Export(answer any) Trace {
	return Trace{
		Problem:   t.problem,
		Algorithm: t.algorithm,
		TreeInput: t.treeInput,
		Nodes:     t.nodes,
		Steps:     t.steps,
		Answer:    answer,
	}
}

// ToJSON serializes the trace to JSON.
func (tr Trace) ToJSON() ([]byte, error) {
	return json.MarshalIndent(tr, "", "  ")
}

// PrintTerminal prints a colored terminal trace (Option 3 experience).
func (tr Trace) PrintTerminal() {
	const (
		reset   = "\033[0m"
		bold    = "\033[1m"
		yellow  = "\033[33m"
		green   = "\033[32m"
		red     = "\033[31m"
		blue    = "\033[34m"
		magenta = "\033[35m"
		cyan    = "\033[36m"
		dim     = "\033[2m"
	)

	fmt.Printf("\n%s%s=== %s — %s ===%s\n", bold, magenta, tr.Problem, tr.Algorithm, reset)
	fmt.Printf("%sInput: %v%s\n\n", dim, tr.TreeInput, reset)

	for _, step := range tr.Steps {
		// Step header
		var actionColor string
		switch step.Action {
		case "visit":
			actionColor = yellow
		case "return":
			actionColor = green
		case "update":
			actionColor = red
		case "init":
			actionColor = blue
		default:
			actionColor = cyan
		}

		fmt.Printf("%s%sStep %d%s · %s%s%s · Node %s\n",
			bold, actionColor, step.ID, reset,
			actionColor, step.Action, reset,
			step.NodeID)

		// Variables
		if len(step.Variables) > 0 {
			for k, v := range step.Variables {
				fmt.Printf("  %s%s%s = %v\n", cyan, k, reset, v)
			}
		}

		// Message
		if step.Message != "" {
			if step.Action == "update" {
				fmt.Printf("  %s%s🏆 %s%s\n", bold, red, step.Message, reset)
			} else {
				fmt.Printf("  %s💬 %s%s\n", dim, step.Message, reset)
			}
		}

		// Call stack
		fmt.Printf("  %sStack: [%s]%s\n\n", dim, strings.Join(step.CallStack, " → "), reset)
	}

	fmt.Printf("%s%s✅ Answer: %v%s\n\n", bold, green, tr.Answer, reset)
}
