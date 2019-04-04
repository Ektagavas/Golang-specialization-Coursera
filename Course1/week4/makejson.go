package main

import(
	"fmt"
	"os"
	"bufio"
	"encoding/json"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var name, add string
	fmt.Printf("Enter name: \n")
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Printf("Enter address: \n")
	if scanner.Scan() {
		add = scanner.Text()
	}

	mp := map[string]string{"name": name, "address": add}
	bytearr, _ := json.Marshal(mp)
	os.Stdout.Write(bytearr)
}

