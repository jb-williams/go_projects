package main

import(
	"fmt"
	"strconv"
	"reflect"
)

var pl = fmt.Println

func main() {
/* casting
	
*/
	// init float
	cV1 := 1.5
	// convert to int, remove decimal part
	cV2 := int(cV1)
	pl(cV2)
	// init a string
	cV3 := "5000000"
	pl(cV3)
	// convert string into interger
	cV4, err := strconv.Atoi(cV3) // Atoi Ascii to Int
	pl(cV4, err, reflect.TypeOf(cV4))
	// string var
	cV5 := 5000000
	// convert int to string
	cV6 := strconv.Itoa(cV5) // Itoa Int to Ascii 
	pl(cV6)
	// handle potential error
	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil{
		pl(cV8)
	}
	// format float into print string
	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)
}
