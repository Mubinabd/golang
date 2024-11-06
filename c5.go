package main

import "fmt"

func main(){
	var burchak1,burchak2,burchak3 int
	fmt.Println("burchak1ni kiriting:")
	fmt.Scan(&burchak1)
	fmt.Println("burchak2ni kiriting:")
	fmt.Scan(&burchak2)
	fmt.Println("burchak3ni kiriting:")
	fmt.Scan(&burchak3)

	if burchak1 > 0 && burchak2 > 0 && burchak3 > 0 && burchak1+burchak2+burchak3 == 180 {
		fmt.Println("Uchburchak")
	} else {
		fmt.Println("Uchburchak emas")
	}
}