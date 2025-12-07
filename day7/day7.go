package main

import (
	"aoc25/util/collections"
	"aoc25/util/grids"
	"aoc25/util/parsing"
	"aoc25/util/point"
	"fmt"
	"log"
)

type Diagram [][]string

func main() {
	diagram := parse()
	solve(diagram)
}

func solve(diagram Diagram) {
	start, ok := grids.FindAny(diagram, "S")

	if !ok {
		log.Fatal("Could not find start")
	}

	visited := collections.NewSet[point.Point]()

	frontier := collections.NewSet[point.Point]()
	frontier.Add(start)

	posMap := make(map[point.Point]int)
	posMap[start] = 1

	splitters := 0

	moved := true
	for moved {
		moved = false

		for _, beam := range frontier.Values() {
			row, col := beam.Y, beam.X

			if !grids.InBounds(diagram, row+1, col) {
				continue
			}

			if diagram[row+1][col] == "^" {
				np := []point.Point{point.Of(col+1, row+1), point.Of(col-1, row+1)}

				splitting := false
				for _, p := range np {
					if grids.InBounds(diagram, p.Y, p.X) {
						frontier.Add(p)
						posMap[p] += posMap[beam]
						moved = true
						splitting = true
					}
				}

				if splitting {
					splitters++
				}
			} else {
				p := point.Of(col, row+1)
				frontier.Add(p)
				moved = true
				posMap[p] += posMap[beam]
			}

			frontier.Remove(beam)
		}

		for _, beam := range frontier.Values() {
			visited.Add(beam)
		}

	}

	total := 0
	for p, count := range posMap {
		if p.Y == len(diagram)-1 {
			total += count
		}
	}

	fmt.Printf("Part A: %d\n", splitters)
	fmt.Printf("Part B: %d\n", total)
}

func parse() Diagram {
	return parsing.FromFile("day7.in").Grid("")
}
