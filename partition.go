package main

import "fmt"

func swap(arr []int, l int ,r int){
    temp := arr[l]
    arr[l] = arr[r]
    arr[r] = temp
    //fmt.Println("Swap Called with l=",l ,"and r=",r)
}

func partition(arr []int, l int, r int) int {
    x := arr[l]
    j := 0
    i:= l+1 
    for ;i < r; i++{
        if arr[i] <= x {
            j = j+1
            swap(arr,j,i)
        }
    }
    swap(arr,l,j)
    return j
}

func quickSort(arr []int, l int, r int){
    if l >= r {
        return
    }
    m := partition(arr,l,r)
    quickSort(arr,l,m-1)
    quickSort(arr,m+1,r)
}

func main(){
     y := []int{4,1,3,6,7,8,9,0} 
     quickSort(y,0,len(y))
     //fmt.Println( partition(y,0,len(y)) )
     fmt.Println(y)
}
