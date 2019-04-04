package main 

import (
	"fmt"
)

type Animal interface{
	Eat()
	Move()
	Speak()
}

// Cow Type
type Cow struct {
	food string
	locomotion string
	noise string
}
func (c Cow) Eat() {
	fmt.Println(c.food)
}
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}
func (c Cow) Speak() {
	fmt.Println(c.noise)
}

// Bird Type
type Bird struct {
	food string
	locomotion string
	noise string
}
func (c Bird) Eat() {
	fmt.Println(c.food)
}
func (c Bird) Move() {
	fmt.Println(c.locomotion)
}
func (c Bird) Speak() {
	fmt.Println(c.noise)
}

// Snake type
type Snake struct {
	food string
	locomotion string
	noise string
}
func (c Snake) Eat() {
	fmt.Println(c.food)
}
func (c Snake) Move() {
	fmt.Println(c.locomotion)
}
func (c Snake) Speak() {
	fmt.Println(c.noise)
}




func main() {
	var request, arg1, arg2 string
	// Map to store created animals
	animals := make(map[string]Animal)
    for {
    	fmt.Printf(">")
		fmt.Scanf("%s %s %s\n",&request, &arg1, &arg2)
		// Create new animal
		if request == "newanimal" {
			switch arg2 {
				case "cow": c := Cow{food: "grass", locomotion: "walk", noise: "moo"}
				animals[arg1] = c
				fmt.Println("Created it!")

				case "bird": b := Bird{food: "worms", locomotion: "fly", noise: "peep"}
				animals[arg1] = b
				fmt.Println("Created it!")

				case "snake": s := Snake{food: "mice", locomotion: "slither", noise: "hsss"}
				animals[arg1] = s
				fmt.Println("Created it!")

				default: fmt.Println("This type of animal is not available")
			}
		} else if request == "query" {
			// Query created animal info
			anim, err := animals[arg1]
			if err == true {
				switch arg2 {
				case "eat": anim.Eat()
				case "move": anim.Move()
				case "speak": anim.Speak()
				default: fmt.Println("Request valid info for the animal")
				} 
			} else {
				fmt.Println("Enter valid name")
			}
			
		} else {
			fmt.Println("Invalid command")
		}
	}
	
}