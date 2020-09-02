package main

import ("fmt"
        "bufio"
        "os"
        "encoding/json"
    )

func main(){
    scanner := bufio.NewScanner(os.Stdin)
    Map := make(map[string]string)
    
    fmt.Println("Enter Your Name :")
    scanner.Scan();
    Map["name"] = scanner.Text();
    fmt.Println("Enter Your Address :")
    scanner.Scan()
    Map["addr"] = scanner.Text()

    barr,err := json.Marshal(Map)
    if err != nil{
        fmt.Println("Error:",err)
    }
    os.Stdout.Write(barr);

}
