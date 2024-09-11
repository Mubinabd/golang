package main

import "fmt"

var (
	lenght = 10
)

func main(){
	s := make([]string,0,lenght)
	for i := range s{
		s[i] = fmt.Sprintf("%d",i+1)
	}
	fmt.Printf("%s",s)
}