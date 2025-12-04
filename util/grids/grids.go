package grids

import (
	"aoc25/util/point"
	"aoc25/util/xmath"
)

func Neighbours8[T any](grid [][]T, row, col int) []point.Point {
	return neighbours(grid, row, col, true)
}

func Neighbours8Values[T any](grid [][]T, row, col int) []T {
	vals := make([]T, 0)
	for _, p := range neighbours(grid, row, col, true) {
		vals = append(vals, grid[p.Y][p.X])
	}

	return vals
}

func Neighbours4[T any](grid [][]T, row, col int) []point.Point {
	return neighbours(grid, row, col, false)
}

func Neighbours4Values[T any](grid [][]T, row, col int) []T {
	vals := make([]T, 0)
	for _, p := range neighbours(grid, row, col, false) {
		vals = append(vals, grid[p.Y][p.X])
	}

	return vals
}

func neighbours[T any](grid [][]T, row, col int, diagonals bool) []point.Point {
	neighbours := []point.Point{}

	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {

			if r == row && c == col {
				continue
			}

			if !diagonals && r != row && c != col {
				continue
			}

			if InBounds(grid, r, c) {
				neighbours = append(neighbours, point.Of(c, r)) // (x,y)
			}
		}
	}

	return neighbours
}

func InBounds[T any](grid [][]T, row, col int) bool {
	if row < 0 || row >= len(grid) {
		return false
	}

	if col < 0 || col >= len(grid[0]) {
		return false
	}

	return true
}

func Map[A, B any](grid [][]A, fn func(A) B) [][]B {
	result := make([][]B, 0, len(grid))
	for _, row := range grid {
		newRow := make([]B, len(row))

		for i, elem := range row {
			newRow[i] = fn(elem)
		}

		result = append(result, newRow)
	}
	return result
}

func RotateRight[A any](grid [][]A, times int) [][]A {

	rotated := grid
	for range times {
		rotated = rotateRight(rotated)
	}

	return rotated
}

func rotateRight[A any](grid [][]A) [][]A {
	if len(grid) == 0 {
		return [][]A{}
	}

	rows := len(grid)
	cols := len(grid[0])

	rotated := make([][]A, cols)
	for i := range rotated {
		rotated[i] = make([]A, rows)
	}

	for r := range rows {
		for c := range cols {
			rotated[c][rows-1-r] = grid[r][c]
		}
	}

	return rotated
}

func Flip[A any](grid [][]A) [][]A {
	rows := len(grid)

	flipped := make([][]A, rows)
	for r := range rows {
		flipped[rows-1-r] = grid[r]
	}

	return flipped
}

func FindAny[A comparable](grid [][]A, a A) (point.Point, bool) {
	for r, row := range grid {
		for c, elem := range row {
			if elem == a {
				return point.Of(c, r), true
			}
		}
	}

	var zero point.Point
	return zero, false
}

func FindAll[A comparable](grid [][]A, a A) []point.Point {
	points := []point.Point{}

	for r, row := range grid {
		for c, elem := range row {
			if elem == a {
				points = append(points, point.Of(r, c))
			}
		}
	}

	return points
}
func ManhattanDistance(a, b point.Point) int {
	return xmath.AbsDiff(a.X, b.X) + xmath.AbsDiff(a.Y, b.Y)
}
