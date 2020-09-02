package main

type Person struct{
    name string
    add string
    phone string
}

func main(){
//    p1 := Person{"Osama","xyz","12345678"};
    p1 := new (Person)
    p1.name ="Osama"
    p1.add  = "George St."
    p1.phone = "12345678"
    p2 := Person{"Huzefa","zxy","998"}
    p2.name = ""
}
