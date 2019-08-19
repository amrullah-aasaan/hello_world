package geometry

import (
	"fmt"
	"math"
)

// Point to represent a point in Cartesian Plane
type Point struct {
	X, Y float64
}

// Distance is used to calculate distance between two cartesian points
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance is used to calculate distance between two cartesian points
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// ScaleBy is used to scale the point by a given factor
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) String() string {
	return fmt.Sprintf("X: %f, Y: %f", p.X, p.Y)
}

// Path is a journey connecting the points with straight lines
type Path []Point

// Distance calculates the total distance covered by this path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
