package main

import ("fmt"
        "sort"
)

func main(){
    //Make a int slice of capacity 3 with length 0
    sli := make([]int,0,3);
    
    //Scan for input, if the input is not an int 
    //then the err will be non-nil
    var input int; 
    _,err := fmt.Scan(&input);
    
    //Loop till the err is nil, i.e the user is entering int
    for err == nil {
        
        //append the input to the slice slice 
        sli = append(sli,input);
        
        //sort the slice and print
        sort.Ints(sli)
        fmt.Println(sli)
        
        //Take input
        _,err = fmt.Scan(&input);
    }
}
