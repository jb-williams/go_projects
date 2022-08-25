package main

import "fmt"

func main() {

	//your code goes here
	i := 4
	fmt.Println(spell(i))
}

func spell(n int) string {
	to10 := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	if n >= 0 && n >= 10 {
		return to10[n]
	}

	return "error"
}
