package main

import "fmt"

func main() {
	var x = []string{"A", "B", "C"}

	for i, s := range x {
		fmt.Println(i, s, " ")
		x = append(x, "Z")
		x[i+1] = "Z"
	}
}
