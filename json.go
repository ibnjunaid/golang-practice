package main 

import "encoding/json"
import "fmt"

type Person struct{
    name string
    addr string
    phone string
}

func main(){
    m := Person{"Osama","2nd St.","123456789"}
    b,err := json.Marshal(m)
    fmt.Println(b)
    fmt.Println(err)
    var k Person
    json.Unmarshal(b,&k)
    fmt.Println(k)
}
