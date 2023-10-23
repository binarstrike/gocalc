package main

import (
	"fmt"

	g "github.com/binarstrike/gocalc/v2"
)

func main() {
	c1 := g.New(20).Multiply(5).Divide(2).Plus(2).Multiply(2).Min(4).Sums()

	fmt.Printf("20 x 5 + 2 x 2 - 4 = %v", c1)
}
