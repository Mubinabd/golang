package main

import "fmt"

func main(){
	defer func ()  {
		fmt.Println(recover())
	}()
	defer func ()  {
		defer func ()  {
			fmt.Println(recover())
		}()
		defer recover()
		panic(2)
	}()
	panic(1)
}
