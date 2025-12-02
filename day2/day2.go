package main

import (
	"aoc25/util/parsing"
	"aoc25/util/point"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	intervals := parse()

	partA(intervals)
	partB(intervals)
}

func partA(intervals []point.Point) {
	invalids := 0

	for _, interval := range intervals {
		start, end := interval.X, interval.Y

		for i := start; i <= end; i++ {
			if invalidA(i) {
				invalids += i
			}
		}
	}

	fmt.Printf("Part A: %d\n", invalids)
}

func partB(intervals []point.Point) {
	invalids := 0

	for _, interval := range intervals {
		start, end := interval.X, interval.Y

		for i := start; i <= end; i++ {
			if invalidB(i) {
				invalids += i
			}
		}
	}

	fmt.Printf("Part B: %d\n", invalids)
}

func invalidA(n int) bool {
	s := strconv.Itoa(n)
	mid := len(s) / 2

	left := s[:mid]
	right := s[mid:]

	return left == right
}

func invalidB(val int) bool {
	s := strconv.Itoa(val)
	n := len(s)

	for size := 1; size < n; size++ {
		if n%size != 0 {
			continue
		}

		head := s[:size]
		allSame := true

		for i := size; i < n; i += size {
			if s[i:i+size] != head {
				allSame = false
				break
			}
		}

		if allSame {
			return true
		}
	}

	return false
}

func parse() []point.Point {

	interval_strs := parsing.FromFile("day2.in").Split(",")
	intervals := make([]point.Point, 0, len(interval_strs))

	for _, interval_str := range interval_strs {
		interval := strings.Split(interval_str, "-")
		start, end := parsing.Stoi(interval[0]), parsing.Stoi(interval[1])
		intervals = append(intervals, point.Of(start, end))
	}

	return intervals
}
