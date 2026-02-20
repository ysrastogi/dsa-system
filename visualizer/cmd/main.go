package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	childrensum "github.com/yashrastogi1/dsa-system/problems/tree/childern_sum_property"
	maxpathsum "github.com/yashrastogi1/dsa-system/problems/tree/max-path-sum"
)

func main() {
	problem := flag.String("problem", "max-path-sum", "Problem to trace")
	input := flag.String("input", "[-10,9,20,null,null,15,7]", "Tree input in LeetCode format")
	outFile := flag.String("out", "", "Output file for trace JSON (default: stdout JSON + terminal trace)")
	flag.Parse()

	// Parse input string to []any
	treeInput := parseInput(*input)

	// Run the traced solution
	var traceResult interface{ ToJSON() ([]byte, error) }

	switch *problem {
	case "max-path-sum":
		t := maxpathsum.MaxPathSumTraced(treeInput)
		t.PrintTerminal()
		traceResult = t
	case "children-sum":
		t := childrensum.ChildrenSumTraced(treeInput)
		t.PrintTerminal()
		traceResult = t
	default:
		fmt.Fprintf(os.Stderr, "Unknown problem: %s
", *problem)
		fmt.Fprintf(os.Stderr, "Available problems: max-path-sum, children-sum
")
		os.Exit(1)
	}

	// Output JSON
	jsonBytes, err := traceResult.ToJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling trace: %v
", err)
		os.Exit(1)
	}

	if *outFile != "" {
		if err := os.WriteFile(*outFile, jsonBytes, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %v
", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Trace written to %s
", *outFile)
	} else {
		fmt.Println(string(jsonBytes))
	}
}

// parseInput converts a LeetCode-format string like "[-10,9,20,null,null,15,7]" to []any.
func parseInput(s string) []any {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "[]")

	if s == "" {
		return nil
	}

	parts := strings.Split(s, ",")
	result := make([]any, len(parts))

	for i, p := range parts {
		p = strings.TrimSpace(p)
		if p == "null" || p == "nil" {
			result[i] = nil
		} else {
			var v int
			if _, err := fmt.Sscanf(p, "%d", &v); err == nil {
				result[i] = v
			} else {
				result[i] = nil
			}
		}
	}

	return result
}

// Trace interface for JSON output
type jsonExporter interface {
	ToJSON() ([]byte, error)
}

// Ensure trace.Trace satisfies the interface at usage site
var _ json.Marshaler = (*json.RawMessage)(nil) // just a compile check placeholder
