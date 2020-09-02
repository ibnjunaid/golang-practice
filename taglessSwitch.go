package main
import "fmt"

// *No need to put break at the end of case statement 

func main(){
        x := 0
        fmt.Println("Enter any Number")
        fmt.Scan(&x)
        switch {
            case x%2 == 0:
                fmt.Println("Even");
            default:
                fmt.Println("Odd");
        }
}
