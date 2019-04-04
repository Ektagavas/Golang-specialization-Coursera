package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	type Fullname struct {
		firstname []byte
		lastname []byte
	}

	sli := make([]Fullname, 0, 1024)

	var filename string
	fmt.Scanln(&filename)

	fd, _:=os.Open(filename)
	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		n := scanner.Text()
		strarr := strings.Split(n, " ")
		st := Fullname{firstname: []byte(strarr[0]), lastname: []byte(strarr[1])}
		sli = append(sli, st)
	}

	fd.Close()

	for _, ele := range sli {
		fmt.Println(string(ele.firstname) + " " + string(ele.lastname))
	}
}

