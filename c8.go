package main

import "fmt"

func main(){
	var son int
	fmt.Println("sonni kiriting:")
	fmt.Scan(&son)

	if son > 0{
		son += 1
	}else if son < 0{
		son -= 2
	}else if son == 0{
		son += 10
	}
	fmt.Println(son)
}