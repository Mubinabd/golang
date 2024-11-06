package main

import "fmt"

func main(){
	var son int
	fmt.Println("sonni kiriting:")
	fmt.Scan(&son)

	if son >= 0{
		son += 1
	}else{
		son -= 2
	}
	fmt.Println(son)
}