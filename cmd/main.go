package main

import (
	"fmt"
	"log"

	patterns "github.com/st-obor/go-patterns/pkg"
)

func main() {

	testOutput(4, 16, -3, 0)
	testOutput(4, 16, -2, 0)
	testOutput(7, 16, -1, 0)
	testOutput(16, 3, 0, 0)
	testOutput(4, 12, 1, 0)
	testOutput(4, 16, 2, 0)
	testOutput(7, 16, 3, 0)
}

func testOutput(n int32, k int32, r int32, g float64) {
	euclid, err := patterns.NewEuclid(n, k, r, g)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n, k, r, g, "-> :\n", euclid.Rhythm)
}
