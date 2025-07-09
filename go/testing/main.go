package main

import (
	"flag"
	"log"
)

func addSomeThing(a, b int) int {
	return a + b
}

func main() {
	a := flag.Int("a", 1, "First # input variabe for addSomeThing().")
	b := flag.Int("b", 2, "Second # input variabe for addSomeThing().")

	sum := addSomeThing(*a, *b)

	log.Printf("%d + %d = %d\n", a, b, sum)
}
