/*
	Goroutines are threads in Go. Interleaved execution of go routines produces race conditions.
	A race condition is that the final value of shared variable depends on the order of execution.
	That is, who gets to update it and when. This results to different results each time you execute 
	the program due to different order of statements execution. The order is undeterministic.

	In the following program, there are two goroutines foo1 and foo2 which share a global variable x.
	Both contain a for loop which iterative increments and decrements x with loop count respectively 
	in foo1 and foo2. When they are executed sequentially, the final value of x is unchanged i.e 10.
	Here, when the two goroutines are executed concurrently, race condition occurs due to shared 
	value of x. One goroutine may read a value of x, and then control may switch to second goroutine 
	which successfully decrements the value of x. When the control again switches back to first
	goroutine, it increments the old value of x it had read previously. Thus, the final value of x
	depends on the order of execution of statements in goroutines.

	To observe the race condition, execute the program more than once and observe the value of x.
	If it still gives same result 10 on your machine, just increase the number of iterations of both
	the for loops to a larger value to observe the race condition.
*/
package main 

import (
	"fmt"
	"time"
)

var x int = 10

func foo1() {
	for i := 0; i < 100000; i++ {
		x = x + i	
	}
}

func foo2() {
	for i := 0; i < 100000; i++ {
		x = x - i	
	}
}

func main() {
	go foo1()
	go foo2()
	time.Sleep(1000) // To prevent main from exiting before goroutines complete
	fmt.Println(x)
}