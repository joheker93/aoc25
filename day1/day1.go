package main

import (
	"aoc25/util/parsing"
	"aoc25/util/xmath"
	"fmt"
)

func main() {
	vals := parse()
	partA(vals)
	partB(vals)
}

func partA(vals []int) {
	ptr := 50
	sum := 0

	for _, val := range vals {
		ptr = (ptr + val) % 100

		if ptr < 0 {
			ptr = 100 + ptr
		}

		if ptr == 0 {
			sum += 1
		}
	}

	fmt.Println(sum)
}

func partB(vals []int) {
	ptr := 50
	sum := 0

	for _, val := range vals {
		dir := 1
		if val < 0 {
			dir = -1
		}

		for i := 0; i < xmath.Abs(val); i++ {
			ptr = (ptr + dir + 100) % 100

			if ptr == 0 {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func parse() []int {
	words := parsing.FromFile("day1.in").Words()
	vals := make([]int, 0, len(words))

	for _, word := range words {
		val := parsing.Stoi(word[1:])

		if word[0] == 'L' {
			val = -val
		}

		vals = append(vals, val)
	}

	return vals
}
