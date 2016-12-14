package main

import (
	"fmt"

	"github.com/bordeltabernacle/go-euler-go/one"
	"github.com/bordeltabernacle/go-euler-go/two"
)

func main() {
	fmt.Printf("\nProblem One: %d", one.SumMultiplesOf(1000, 3, 5))
	fmt.Printf("\nProblem Two: %d", two.SumEvenFib(4e6))
}
