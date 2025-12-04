package main

import (
	"aoc25/util/grids"
	"aoc25/util/parsing"
	"aoc25/util/xslices"
	"fmt"
)

type Grid [][]string

func main() {
	grid := parse()
	partA(grid)
	partB(grid)
}

func partA(grid Grid) {
	accessables := 0

	for r, row := range grid {
		for c := range row {
			if accessable(grid, r, c) {
				accessables++
			}
		}
	}

	fmt.Printf("Part A: %d", accessables)
}

func partB(grid Grid) {
	removed := 0

	for {
		couldRemove := false
		for r, row := range grid {
			for c := range row {
				if !accessable(grid, r, c) {
					continue
				}

				removeRoll(grid, r, c)
				couldRemove = true
				removed++
			}
		}

		if !couldRemove {
			break
		}
	}

	fmt.Printf("Part B: %d\n", removed)
}

func accessable(grid Grid, row, col int) bool {
	if grid[row][col] == "@" {
		neighbours := grids.Neighbours8Values(grid, row, col)
		rolls := xslices.Frequency("@", neighbours)

		if rolls < 4 {
			return true
		}
	}

	return false
}

func removeRoll(grid Grid, row, col int) {
	grid[row][col] = "."
}

func parse() Grid {
	return parsing.FromFile("day4.in").Grid("")
}
