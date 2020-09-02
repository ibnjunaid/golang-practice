package main

import "fmt"

func echo() int {
    fmt.Println("Hello")
    return 4
}

func main(){
    i:=1
    defer fmt.Println(echo())
    i++;
    fmt.Println(i)
}
