package main 

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var sem = make(chan int, 2)


type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	leftCS, rightCS *Chopstick
	id int
}

func (p Philosopher) Eat(h *Host) {
	for i:=0; i<3; i++ {
		if h.Acquire() == 1 {
			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Println("Starting to eat", p.id)

			fmt.Println("Finishing eating", p.id)

			p.rightCS.Unlock()
			p.leftCS.Unlock()

			h.Release()
			
		}
	}
	wg.Done()
}


type Host struct {
}

func (h Host) Acquire() int{
	sem <- 1
	return 1
}

func (h Host) Release() {
	<-sem
}


func main() {

	var h Host

	CS := make([] *Chopstick, 5)
	for i := 0; i < 5; i++ {
		CS[i] = new(Chopstick)
	}

	Philos := make([] *Philosopher, 5)
	for i:=0; i<5; i++ {
		Philos[i] = &Philosopher{leftCS: CS[i], rightCS: CS[(i+1)%5], id: i+1}		
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go Philos[i].Eat(&h)
	}

	wg.Wait()
}


