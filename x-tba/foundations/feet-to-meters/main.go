package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	c, _ := strconv.ParseFloat(os.Args[1], 64)
	f := c*1.8 + 32

	// Like this:
	fmt.Printf("%g ºC is %g ºF\n", c, f)

	// Or just like this (both are correct):
	fmt.Printf("%g ºF\n", f)
}
