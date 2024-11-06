package main

import "fmt"

func main() {
    var son int
	fmt.Printf("Son kiriting: ")
	fmt.Scan(&son)

	if son  >= 0{
		son += 1
	}else{
		fmt.Println(son)
	}
	fmt.Println(son)
}