package main

import (
        "fmt"
        "strings"
        "bufio"
        "os"
       )
func check(input string)string{
    input = strings.ToLower(input)  
    length := len(input) - 1
    if(input[0] == 'i' && input[length] == 'n'){
        for i:=1; i<length ;i++{
            if(input[i] == 'a'){
                return "Found"
            }
        }
        return "Not Found"
    }
    return "Not Found"
}

func main(){
    //used bufio for having whitespaces in the input string
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    fmt.Println(check(line))
}
