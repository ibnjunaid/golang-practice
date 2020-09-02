package main

import (
        "encoding/json"
        "fmt"
        "os"
    )

func main(){
    arr := "My string"
    dat,_:= json.Marshal(arr)
    fmt.Println(dat)
    os.Stdout.Write(dat)
}
