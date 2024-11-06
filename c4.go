package main

import "fmt"

func main() {
    var son int
    fmt.Print("Son kiriting: ")
    fmt.Scan(&son) 

    if son % 3 == 0 {
        fmt.Println("bo'linadi")
    } else {
        fmt.Println("bo'linmaydi")
    }
}
