package main 

import (
	"fmt"
)


func GenDisplaceFn(a,vo,so float64) func(float64) float64{
	// Function to calculate displacement given time
	funcVar := func (t float64) float64{
		return ((0.5 * a * t * t) + (vo * t) + so)
	}
	return funcVar
}

func main() {
	var a, vo, so, t float64

	fmt.Printf("Enter acceleration:\n")
	fmt.Scanln(&a)
	fmt.Printf("Enter initial velocity:\n")
	fmt.Scanln(&vo)
	fmt.Printf("Enter initial displacement:\n")
	fmt.Scanln(&so)
	fmt.Printf("Enter time:\n")
	fmt.Scanln(&t)

	fn := GenDisplaceFn(a, vo, so)
	fmt.Println(fn(t))

}