package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x / 2)
	//z := 1.0
	prevZ := z
	i := 1
	for i <= 100 {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prevZ) < 0.000001 {
			break
		}
		prevZ = z
		i += 1
	}
	fmt.Println(i)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))

	fmt.Println(Sqrt(234282402323241))
	fmt.Println(math.Sqrt(234282402323241))
	/*fmt.Println(Sqrt(3))
	  fmt.Println(Sqrt(4))
	  fmt.Println(Sqrt(5))*/
}
