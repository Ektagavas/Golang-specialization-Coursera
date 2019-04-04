package main 

import (
	"fmt"
	"sync"
)


// Bubble sort
func BubbleSort(sli []int, wg *sync.WaitGroup) {
	fmt.Println(sli)
	for i:=0; i<(len(sli)-1); i++ {
		for j:=0; j<(len(sli)-1); j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
	wg.Done()
}


// Swap the consecutive elements of slice at given index
func Swap(sli []int, i int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}

// Partition input intp four parts and return indices of parts
func Partition(n, parts int) []int {
	sli := make([]int, 0, parts)

	eachpart := int(n / parts)
	j:=0
	for i:=0 ; i < parts; i++{
		sli = append(sli, j)
		j += eachpart
	}
	return sli
}


// Merge two sorted slices into one sorted slice
func Merge(sli []int, index1, index2, end int) []int {
	newsli := make([]int,0, end-index1)
	i := index1
	j := index2

	for ; (i < index2) && (j < end); {
		if(sli[i] < sli[j]) {
			newsli = append(newsli, sli[i])
			i++
		} else {
			newsli = append(newsli, sli[j])
			j++
		}
	}

	if(i == index2 && j < end) {
		newsli = append(newsli, sli[j:end]...)
	} else if (j == end && i<index2) {
		newsli = append(newsli, sli[i:index2]...)
	}

	return newsli
}


func main() {
	var n, parts int
	parts = 4
	var wg sync.WaitGroup

	fmt.Print("Enter number of integers to sort (>= 4): ")
	fmt.Scanln(&n)

	for ; n<4; {
		fmt.Print("Enter number of integers to sort (>= 4): ")
		fmt.Scanln(&n)
	}

	fmt.Println("Enter elements to sort separated by new line: ")
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var g int
		fmt.Scanln(&g)
		arr = append(arr, g)
	}

	sli := Partition(n, parts)
	wg.Add(parts)
	sli = append(sli, n)

	// Go routines equal to required parts
	for i := 0; i < len(sli)-1; i++ {
		go BubbleSort(arr[sli[i]:sli[i+1]], &wg)
	}

	wg.Wait()

	// Merge
	for i := 0; i < len(sli)-2; i++ {
			tarr := Merge(arr, 0, sli[i+1], sli[i+2])
			for k:=0; k<len(tarr); k++ {
				arr[k] = tarr[k]
			}
	}

	fmt.Println("Final array", arr)

}