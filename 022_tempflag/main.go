package main

import (
	"flag"
	"fmt"

	"github.com/amrullah-aasaan/hello_world/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
