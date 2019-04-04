package main 

import (
	"fmt"
)

type Animal struct {
		food string
		locomotion string
		noise string
	}

func (a *Animal) Eat() {
	fmt.Printf("%s\n", a.food)
}

func (a *Animal) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Printf("%s\n", a.noise) 
}

func main() {
	

	animals := map[string] Animal {"cow": Animal{food: "grass", locomotion: "walk", noise: "moo"},
								   "bird": Animal{food: "worms", locomotion: "fly", noise: "peep"},
								   "snake": Animal{food: "mice", locomotion: "slither", noise: "hsss"}}
	
    var name, info string
    for {
    	fmt.Printf(">")
		fmt.Scanf("%s %s\n",&name, &info)
		anim := animals[name]
		switch info {
			case "eat": anim.Eat()
			case "speak": anim.Speak()
			case "move": anim.Move()
		}
    }
	
}

