package main

import(
	"fmt"
	"bufio"
	"os"
	"log"
)

var pl = fmt.Println

func main() {
// gets user name and print it
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("hello", name)
	} else {
		log.Fatal(err)
	}
}
