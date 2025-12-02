package xmath

import (
	"golang.org/x/exp/constraints"
)

func AbsDiff[Num constraints.Integer | constraints.Float](a, b Num) Num {
	if a >= b {
		return a - b
	}

	return b - a
}

func Abs[Num constraints.Integer | constraints.Float](a Num) Num {
	if a < 0 {
		return -a
	}

	return a
}

func SumOfDigits[Int constraints.Integer](a Int) Int {

	var sum Int

	for a >= 10 {
		sum += a % 10
		a = a / 10
	}

	return sum + a%10
}

func GCD[Int constraints.Integer](a, b Int) Int {
	var gcd Int

	for a%b != 0 {
		gcd = a % b
		a = b
		b = gcd
	}

	return gcd
}
