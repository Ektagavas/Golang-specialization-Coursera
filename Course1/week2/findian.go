package main
import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s string
	fmt.Printf("Enter a string\n")
	if scanner.Scan() {
		s = scanner.Text()     // Read string with spaces
	}

	s = strings.ToLower(s)     // To make case insensitive

	if strings.HasPrefix(s,"i") && strings.HasSuffix(s,"n") && strings.Contains(s,"a") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}

}
