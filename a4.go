package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	for i, v := range x {
		if v == "b" {
			x = append(x[:i],
				x[i+1:]...)
		} else {
			fmt.Println(v, " ")
		}
	}
}