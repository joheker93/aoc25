package point

type Point struct {
	X, Y int
}

func Of(x, y int) Point {
	return Point{x, y}
}

func (point Point) Unpack() (x, y int) {
	return point.X, point.Y
}

func (point Point) Swap() Point {
	return Point{X: point.Y, Y: point.X}
}
