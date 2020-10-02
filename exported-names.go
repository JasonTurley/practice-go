package main

import (
	"fmt"
	"math"
)

func main() {
	// gives error because "pi" is an unexported-name
	// fmt.Println(math.pi)

	// works because "Pi" is an exported-name
	fmt.Println(math.Pi)
}
