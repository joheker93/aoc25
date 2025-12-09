package main

import (
	"aoc25/util/parsing"
	"aoc25/util/xslices"
	"fmt"
	"maps"
	"math"
	"slices"
	"sort"
	"strings"
)

type Junction struct {
	x, y, z int
}

type Distance struct {
	a, b int // indexes
	dist float64
}

func main() {
	junctions := parse()
	solve(junctions)
}

func solve(junctions []Junction) {

	distances := make([]Distance, 0)
	for i, j1 := range junctions {
		for k, j2 := range junctions[i+1:] {
			i2 := i + k + 1
			distances = append(distances, Distance{a: i, b: i2, dist: distance(j1, j2)})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist <= distances[j].dist
	})

	circuits := make(map[Junction]int)

	for i, junction := range junctions {
		circuits[junction] = i
	}

	for i, distance := range distances {
		j1, j2 := junctions[distance.a], junctions[distance.b]

		j1Circuit := circuits[j1]
		j2Circuit := circuits[j2]
		for _, junction := range junctions {
			if circuits[junction] == j2Circuit {
				circuits[junction] = j1Circuit
			}
		}

		if i == 1000 {
			solveA(circuits)
		}

		solved := solveB(circuits, junctions, distance)
		if solved {
			break
		}
	}

}

func solveA(circuits map[Junction]int) {
	counts := make(map[int]int)
	for _, circuit := range circuits {
		counts[circuit]++
	}

	countSlice := slices.Collect(maps.Values(counts))
	sort.Slice(countSlice, func(i, j int) bool {
		return countSlice[i] >= countSlice[j]
	})

	result := xslices.Product(countSlice[:3])

	fmt.Printf("Part A: %d\n", result)
}

func solveB(circuits map[Junction]int, junctions []Junction, distance Distance) bool {
	allMatch := true
	circuit := circuits[junctions[0]]
	for _, junction := range junctions[1:] {
		if circuits[junction] != circuit {
			allMatch = false
		}
	}

	if allMatch {
		j1, j2 := junctions[distance.a], junctions[distance.b]
		fmt.Printf("Part B: %d\n", j1.x*j2.x)
		return true
	}

	return false
}

func distance(a, b Junction) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	dz := float64(a.z - b.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func parse() []Junction {
	lines := parsing.FromFile("day8.in").Lines().Get()

	junctions := make([]Junction, 0, len(lines))
	for _, line := range lines {
		xyz := strings.Split(line, ",")
		junctions = append(junctions, Junction{x: parsing.Stoi(xyz[0]), y: parsing.Stoi(xyz[1]), z: parsing.Stoi(xyz[2])})
	}

	return junctions
}
