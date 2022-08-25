package main

import(
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
/* core data types: int float64, bool, string, rune
Default type 0, 0.0, false, ""
*/
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf("Hello"))

}
