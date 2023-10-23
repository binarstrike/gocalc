package main

import (
	"fmt"

	g "github.com/binarstrike/gocalc"
)

func main() {
	c1 := g.NewCalc(10).Multiply(10).Multiply(10).Divide(500).Sums()

	fmt.Printf("10 x 10 x 10 / 500 = %v", c1)
}
