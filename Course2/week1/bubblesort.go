package main 
import (
	"fmt"
)

// Bubble sort
func BubbleSort(sli []int) {
	for i:=0; i<(len(sli)-1); i++ {
		for j:=0; j<(len(sli)-1); j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}


// Swap the consecutive elements of slice at given index
func Swap(sli []int, i int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}


func main() {
	arr := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Printf("\nEnter %d integer: ", (i+1))
		fmt.Scanln(&arr[i])
	}

	BubbleSort(arr[0:len(arr)])

	fmt.Println(arr)
}