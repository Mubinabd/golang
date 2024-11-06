package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Sonlarni kiriting:")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	var count int
	if a > 0 {
		count++
	}
	if b > 0 {
		count++
	}
	if c > 0 {
		count++
	}

	if count > 0 {
		fmt.Printf("Musbat sonlar soni: %d\n", count)
	} else {
		fmt.Println("Musbat son topilmadi")
	}
}
