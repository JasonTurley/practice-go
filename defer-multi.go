// Deferred function calls are pushed onto the stack and executed in a
// last-in-first-out order.
// Learn more about defer statements here:
// https://blog.golang.org/defer-panic-and-recover

package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
	defer fmt.Println(i)
    }

    fmt.Println("done")
}
