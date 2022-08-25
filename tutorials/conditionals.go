package main
// works with errors got to figure that out sometime

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
/* Conditional Operators: > < >= <= == !=
   Logical Operators : && || !
*/
   pl("What is your age?")
   reader := bufio.NewReader(os.Stdin)
   iAgeS, err := reader.ReadString('\n')
   iAgeI, err := strconv.Atoi(iAgeS)
   if (iAgeI >= 1) && (iAgeI <= 18){
	   pl("Important Birthday", err)
   } else if (iAgeI == 21) || (iAgeI == 50){
	   pl("Important Birthday", err)
   } else if iAgeI >= 65 {
	   pl("Important Birthday", err)
   } else {
	   pl("Not an Important Birthday", err)
   }
   pl("!true =", !true)
}
