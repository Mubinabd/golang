package main

import "fmt"

var X int

func init() { X = 1 }

var x2 = 2 * X

func main() {
	fmt.Println(x2)
}
