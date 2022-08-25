package main

import(
	"fmt"
	"time"
	"math"
	"math/rand"
)

var pl = fmt.Println

func main() {
	pl("5 + 4 =", 5+4)
	pl("5 - 4 =", 5-4)
	pl("5 * 4 =", 5*4)
	pl("5 / 4 =", 5/4)
	pl("5 % 4 =", 5%4)
	mInt := 1
	pl(mInt)
	mInt = mInt + 1
	pl(mInt)
	mInt += 1
	pl(mInt)
	mInt++
	pl(mInt)
	pl("Float Precision =", 0.111111111111111111 + 0.111111111111111111)
	// random values
	// cool to get seed value for a random number generator using time(seconds, seconds since the date 1/1/70)
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)
// There are many math functions
	pl("Abs(-10) =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4, 2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	// Get the log of e to the power of 2
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5,4) =", math.Max(5, 4))
	pl("Min(5,4) =", math.Min(5, 4))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	// Convert 1.5708 radians to degrees
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)
	fmt.Printf("%f radians = %.2f degrees\n", r90, d90)
	fmt.Printf("%.0f radians = %.0f degrees\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))
}
