package main
import(
	"fmt"
	"strconv"
	"sort"
)

func main() {
	sli := make([]int, 0, 3)

	for {
		var i string
		fmt.Printf("\nEnter an integer: ")
		fmt.Scan(&i)
		if i == "X"{
			// Exit the loop
			break
		} else {
			// Convert string to int
			j, err := strconv.Atoi(i)
			if err == nil {
				// Append to slice
				sli = append(sli, j)

				// Sort and print the slice
				sort.Ints(sli)
				fmt.Println(sli)
			}
			
		}
	}
}