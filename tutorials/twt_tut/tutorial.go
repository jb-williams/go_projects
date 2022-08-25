package main

import(
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	pl("Welcome to my quiz game!")

	var name string = "Tim"
	name = "hello"
	pl(name)
}

