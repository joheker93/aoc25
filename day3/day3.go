package main

import (
	"aoc25/util/parsing"
	"fmt"
)

type Banks [][]int

func main() {
	banks := parse()
	partA(banks)
	partB(banks)
}

func partA(banks Banks) {
	joltage := computeJoltage(banks, 2)
	fmt.Printf("Part A: %d\n", joltage)
}

func partB(banks Banks) {
	joltage := computeJoltage(banks, 12)
	fmt.Printf("Part B: %d\n", joltage)
}

func computeJoltage(banks Banks, cells int) int {
	joltage := 0
	for _, bank := range banks {
		joltage += highestJoltage(bank, cells)
	}

	return joltage
}

func highestJoltage(bank []int, cells int) int {
	start := 0
	stop := len(bank) - cells + 1
	maxes := make([]int, 0)

	for len(maxes) < cells {

		currentMax := 0
		currentMaxIdx := 0

		for i := start; i < stop; i++ {
			if bank[i] > currentMax {
				currentMax = bank[i]
				currentMaxIdx = i
			}
		}

		start = currentMaxIdx + 1
		stop++

		maxes = append(maxes, currentMax)
	}

	n := 0
	for _, d := range maxes {
		n = n*10 + d
	}
	return n
}

func parse() [][]int {
	lines := parsing.FromFile("day3.in").Lines().Get()

	banks := make([][]int, 0)
	for _, line := range lines {
		bank := make([]int, 0)
		for _, ch := range line {
			bank = append(bank, int(ch-'0'))
		}

		banks = append(banks, bank)
	}

	return banks
}
