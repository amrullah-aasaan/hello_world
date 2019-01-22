package main

import (
	"fmt"
	"os"
	"strings"
)

// This is a comment
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
