package main

import (
	"fmt"

	"github.com/binarstrike/gocalc/v2"
)

func main() {
	c := gocalc.New(10)
	c.Plus(10)
	c.Multiply(5)

	fmt.Println(c.Result())
}
