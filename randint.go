// The environment is deterministic, so each time this program is run it
// will return the same number.
// To see a different number use rand.Seed
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("My favorite number is", rand.Intn(10))
}
