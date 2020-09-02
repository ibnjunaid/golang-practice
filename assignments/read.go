package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	fname string
	lname string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter File Name:")
	scanner.Scan()
	filename := scanner.Text()
	file, err := os.Open(filename)
	checkError(err)
	Pslice := make([]Person, 0, 10)
    data := make([]byte, 40) 
	for {
		n, _ := file.Read(data)
		if n == 0 {
			break
		}
        var localPerson Person
        localPerson.fname = string(data[:20])
        localPerson.lname = string(data[21:40])
        Pslice = append(Pslice,localPerson)
    }
    fmt.Println(len(Pslice))
    for i := 0; i < len(Pslice); i++ {
        fmt.Println("FirstName of Person: ",Pslice[i].fname)
        fmt.Println("lastName of Person: ",Pslice[i].lname)
    }
}