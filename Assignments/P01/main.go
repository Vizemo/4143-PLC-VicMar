package main

import (
	"fmt"

	"example.com/go-demo-1/mascot"

	"rsc.io/quote"
)

// function that prints out the best mascot and a quote from the quote library
func main() {
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
}
