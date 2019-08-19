package main

import (
	"fmt"

	"github.com/amrullah-aasaan/hello_world/geometry"
)

func main() {
	perim := geometry.Path{
		{X: 1, Y: 1}, {X: 5, Y: 1}, {X: 5, Y: 4}, {X: 1, Y: 1},
	}
	fmt.Println("Total Distance covered by points:", perim.Distance())

	// various ways to call a pointer method
	p1 := &geometry.Point{X: 1, Y: 2}
	p1.ScaleBy(2)
	fmt.Println("Scaled Point 1:", *p1)

	p2 := geometry.Point{X: 1, Y: 2}
	(&p2).ScaleBy(2)
	fmt.Println("Scaled Point 2:", p2)

	p3 := geometry.Point{X: 1, Y: 2}
	p3.ScaleBy(2)  // convenience offered by golang
	fmt.Println("Scaled Point 3:", p3)

}
