package main

import "fmt"

func swap(arr []int , index int){
    temp := arr[index+1]
    arr[index+1] = arr[index]
    arr[index] = temp
}

func bubbleSort(array []int){
    for i:=0; i<len(array)-1; i++{
       if array[i] > array[i+1]{
            swap(array,i)
       }
    }
}

func main(){
    array := []int{6,4,8,2,9,3,9,4,7,6,1}
    fmt.Println(array)
    bubbleSort(array)
    fmt.Println(array)
}
