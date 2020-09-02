package main

import "fmt"

func main(){
    var idMap map[string]int
    idMap = make(map[string]int)
    for i:=32; i<50; i++{
        idMap[string(i)] = i
    }

    for key,val := range idMap{
        fmt.Print(key);
        fmt.Print(":");
        fmt.Print(val);
        fmt.Println();
    }
}
