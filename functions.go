package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func diff(x int, y int) int {
	return x - y
}

func product(x int, y int) int {
	return x * y
}

func main() {
	x := 4
	y := 2

	fmt.Printf("%d + %d = %d\n", x, y, sum(x, y))
	fmt.Printf("%d - %d = %d\n", x, y, diff(x, y))
	fmt.Printf("%d * %d = %d\n", x, y, product(x, y))
}
