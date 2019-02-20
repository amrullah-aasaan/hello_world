package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
	w2 := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // Trailing comma is necessary
	}
	fmt.Printf("%#v\n", w1)
	w2.X = 42 // directly accessible without doing w2.Circle.Point.X
	fmt.Printf("%#v\n", w2)
}
