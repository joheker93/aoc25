package main

import (
	"aoc25/util/grids"
	"aoc25/util/parsing"
	"aoc25/util/xslices"
	"fmt"
	"strings"
)

func main() {
	equations := parseA()
	partA(equations)

	eqs, ops := parseB()
	partB(eqs, ops)

}

func partA(equations [][]string) {
	total := 0
	for _, equations := range equations {
		localResult := parsing.Stoi(equations[0])
		end := len(equations)
		op := equations[end-1]
		for _, elem := range equations[1 : end-1] {
			if op == "+" {
				localResult += parsing.Stoi(elem)
				continue
			}

			if op == "*" {
				localResult *= parsing.Stoi(elem)
				continue
			}
		}

		total += localResult
	}

	fmt.Printf("Part A: %d\n", total)

}

func partB(eqs [][]string, ops []string) {
	total := 0
	localValues := make([]int, 0)
	op := 0

	applyOp := func(values []int, operator string) int {
		switch operator {
		case "*":
			return xslices.Product(values)
		case "+":
			return xslices.Sum(values)
		default:
			return 0
		}
	}

	for _, eq := range eqs {
		if len(eq) == 0 {
			if op < len(ops) {
				total += applyOp(localValues, ops[op])
				op++
			}
			localValues = nil
			continue
		}

		val := parsing.Stoi(strings.Join(eq, ""))
		localValues = append(localValues, val)
	}

	if len(localValues) > 0 && op < len(ops) {
		total += applyOp(localValues, ops[op])
	}

	fmt.Printf("Part B: %d\n", total)
}

func parseA() [][]string {
	lines := parsing.FromFile("day6.in").Lines()
	equations := make([][]string, 0)
	for _, line := range lines {
		words := parsing.FromString(line).Words().Get()
		equations = append(equations, words)
	}

	return grids.Transpose(equations)
}

func parseB() ([][]string, []string) {
	lines := parsing.FromFile("day6.in").Lines().Get()
	newLines := make([][]string, 0)

	for _, line := range lines {
		strs := make([]string, 0)
		for _, b := range line {
			str := string(b)
			strs = append(strs, str)
		}

		newLines = append(newLines, strs)
	}

	ops := newLines[len(newLines)-1]
	newLines = grids.Transpose(newLines[:len(newLines)-1])

	for i, line := range newLines {
		filtered := line[:0]
		for _, elem := range line {
			if strings.TrimSpace(elem) != "" {
				filtered = append(filtered, elem)
			}
		}
		newLines[i] = filtered
	}

	filtOps := make([]string, 0)

	for _, op := range ops {
		if strings.TrimSpace(op) != "" {
			filtOps = append(filtOps, op)
		}
	}

	return newLines, filtOps
}
