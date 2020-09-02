package main

import "fmt"

func sum(arr [4]int) int {
    var s int = 0
    for _, elem := range arr{
        s += elem
    }
    return s
}

func sumbyRef(arr *[4]int) int {
    sum := 0
    for _,elem := range arr{
        sum += elem
    }
    return sum;
}

func sumByslice(sli []int) int{
    sum := 0
    length := len(sli)
    for i:=0; i<length ;i++{
        sum += sli[i];
    }
    return sum
}
func main(){
    val := []int{1,2,3,4};
    //ret:= sum(val)
    //fmt.Println(ret);
    fmt.Println( sumByslice(val) )
    //fmt.Println( sumbyRef(&val) )
}

