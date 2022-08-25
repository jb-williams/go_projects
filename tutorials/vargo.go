package main

import(
	"fmt"
)

var pl = fmt.Println

func main() {
/*  var name type
	begin with letter, contain letter and/or number
	Starts with capital var is exported and can be access outside the package
	camelCase is standard naming convention
	if using := only works for the initialization of the var
*/
	var vName string = "Jb"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 2.4
	pl(vName)
	pl(v1)
	pl(v2)
	pl(v3)
	pl(v4)

}
