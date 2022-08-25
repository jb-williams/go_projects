package main

import "fmt"

func main() {
	a := 20
	fmt.Println(a)

	// gives error if already declared
	a := 30
	fmt.Println(a)
}
