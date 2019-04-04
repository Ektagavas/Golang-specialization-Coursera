package main
import "fmt"

func main() {
	var n float32
	fmt.Printf("Enter a floating point number\n")
	fmt.Scan(&n)
	var in int = int(n)   // Truncate to int
	fmt.Printf("Truncated number: %d", in)	
}
