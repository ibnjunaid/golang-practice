package main
import "fmt"

// *No need to put break at the end of case statement 

func main(){
    x := 10
    switch x {
        case 1:
            fmt.Println(1);
        case 2:
            fmt.Println(2);
        default:
            fmt.Println(x);
    }
}

