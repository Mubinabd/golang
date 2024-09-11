package main

import "fmt"

func main(){
	s := []int{0,1,2,3,4}
	copy(s[2:],s[0:])
	fmt.Println(s)
}