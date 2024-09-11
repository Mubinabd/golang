package main

import "fmt"

func sum(nums ...float64) float64 {
	s := 0.0
	for _, v := range nums {
		s += v
	}

	return s
}
func main() {
	fmt.Println(sum(0.1, 0.2, 03) == sum(0.3, 0.2, 0.1))
}
