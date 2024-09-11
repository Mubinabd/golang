package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	for i, v := range x {
		i++
		fmt.Print(i, v, ",")
	}
}
