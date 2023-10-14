package main

import "fmt"

func main() {
	sum := 1

	// gofmt is opinionated
	// "Gofmt's style is no one's favorite, yet gofmt is everyone's favorite"
	// and is always removing the semi-colons
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
