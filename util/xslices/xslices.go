package xslices

import (
	"aoc25/util/pair"
	"cmp"
	"errors"

	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Integer | constraints.Float](ts []T) T {
	var sum T
	for _, t := range ts {
		sum += t
	}

	return sum
}

func Product[T constraints.Integer | constraints.Float](ts []T) T {
	var prod T = 1
	for _, t := range ts {
		prod *= t
	}

	return prod
}

func Map[A, B any](as []A, fn func(A) B) []B {
	bs := make([]B, len(as))

	for i, a := range as {
		bs[i] = fn(a)
	}

	return bs
}

func Map2[A, B any](as [][]A, fn func(A) B) [][]B {
	bs := make([][]B, len(as))
	for i := range as {
		bs[i] = Map(as[i], fn)
	}

	return bs
}

func Fold[A, B any](fn func(B, A) B, acc B, as []A) B {
	for _, a := range as {
		acc = fn(acc, a)
	}
	return acc
}

func Frequency[A comparable](match A, as []A) int {
	counts := make(map[A]int)

	for _, a := range as {
		counts[a] += 1
	}

	return counts[match]
}

func Frequencies[A comparable](as []A) map[A]int {
	counts := make(map[A]int)
	for _, a := range as {
		counts[a] += 1
	}

	return counts
}

func Zip[A, B any](as []A, bs []B) []pair.Pair[A, B] {
	pairs := make([]pair.Pair[A, B], len(as))
	for i := range as {
		pairs[i] = pair.Of(as[i], bs[i])
	}

	return pairs
}

func FindAny[A any](as []A, predicate func(A) bool) (A, error) {
	for _, a := range as {
		if predicate(a) {
			return a, nil
		}
	}

	var a A
	return a, errors.New("no element matched the predicate")
}

func Min[A cmp.Ordered](as []A) (A, error) {
	return minMax(as, func(a, b A) bool { return a <= b })
}

func Max[A cmp.Ordered](as []A) (A, error) {
	return minMax(as, func(a, b A) bool { return a >= b })
}

func minMax[A cmp.Ordered](as []A, predicate func(a, b A) bool) (A, error) {
	if len(as) == 0 {
		var zero A
		return zero, errors.New("empty list")
	}

	max := as[0]

	for _, a := range as {
		if predicate(a, max) {
			max = a
		}
	}

	return max, nil
}
