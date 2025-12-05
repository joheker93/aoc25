package main

import (
	"aoc25/util/parsing"
	"aoc25/util/point"
	"fmt"
	"sort"
	"strings"
)

type DataBase struct {
	freshRanges []point.Point
	ids         []int
}

func main() {
	db := parse()
	partA(db)
	partB(db)
}

func partA(db DataBase) {

	fresh := 0
	for _, id := range db.ids {
		if isFresh(id, db) {
			fresh++
		}
	}

	fmt.Printf("Part A: %d\n", fresh)
}

func isFresh(id int, db DataBase) bool {
	for _, interval := range db.freshRanges {
		start, end := interval.X, interval.Y

		if id >= start && id <= end {
			return true
		}
	}

	return false
}

func partB(db DataBase) {

	sort.Slice(db.freshRanges, func(i, j int) bool {
		return db.freshRanges[i].X < db.freshRanges[j].X
	})

	merged := make([]point.Point, 0)
	start, end := db.freshRanges[0].X, db.freshRanges[0].Y

	for _, interval := range db.freshRanges[1:] {
		next_start, next_end := interval.X, interval.Y

		if next_start <= end+1 {
			end = max(end, next_end)
		} else {
			merged = append(merged, point.Of(start, end))
			start, end = next_start, next_end
		}

	}

	merged = append(merged, point.Of(start, end))

	total := 0

	for _, m := range merged {
		total += (m.Y - m.X + 1)
	}

	fmt.Printf("Part B: %d\n", total)
}

func parse() DataBase {
	lines := parsing.FromFile("day5.in").Lines().Get()

	var db DataBase = DataBase{freshRanges: make([]point.Point, 0), ids: make([]int, 0)}

	intervals := true

	for _, line := range lines {
		if line == "" {
			intervals = false
			continue
		}

		if intervals {
			interval := strings.Split(line, "-")
			db.freshRanges = append(db.freshRanges, point.Of(parsing.Stoi(interval[0]), parsing.Stoi(interval[1])))
		} else {
			db.ids = append(db.ids, parsing.Stoi(line))
		}
	}

	return db
}
