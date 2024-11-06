package main

import "fmt"

// func main() {
// 	var a, b, c int
// 	fmt.Println("Sonlarni kiriting:")
// 	fmt.Scan(&a)
// 	fmt.Scan(&b)
// 	fmt.Scan(&c)

// 	var count int
// 	var count2 int
// 	if a > 0 {
// 		count++
// 	}else{
// 		count2++
// 	}
// 	if b > 0 {
// 		count++
// 	}else{
// 		count2++
// 	}
// 	if c > 0 {
// 		count++
// 	}else{
// 		count2++
// 	}

// 	if count > 0 && count2 > 0{
// 		fmt.Printf("Musbat sonlar soni: %d\n", count)
// 		fmt.Printf("Manfiy sonlar soni: %d\n", count2)
// 	}
// }

// func main(){
// 	var a,b int
// 	fmt.Println("son kirit:")
// 	fmt.Scan(&a)
// 	fmt.Scan(&b)

// 	if a > b {
// 		fmt.Printf("a son katta:",a)
// 	}else{
// 		fmt.Printf("b son katta:",b)
// 	}
// }

// func main() {
// 	var a, b int
// 	fmt.Println("son kirit:")
// 	fmt.Scan(&a)
// 	fmt.Scan(&b)

// 	if a < b {
// 		fmt.Println("a son kichik:", a)
// 	} else if a > b{
// 		fmt.Println("b son kichik:", b)
// 	}
// }

// func main() {
// 	var a, b int
// 	fmt.Println("son kirit:")
// 	fmt.Scan(&a)
// 	fmt.Scan(&b)

// 	if a > b || a < b{
// 		fmt.Println("b son kichik:", b)
// 		fmt.Println("a son katta:", a)
// 	}
// }


// func main() {
// 	var a, b int
// 	fmt.Println("Sonlarni kiriting:")
// 	fmt.Scan(&a)
// 	fmt.Scan(&b)

// 	if a < b {
// 		a, b = b, a
// 	}

// 	fmt.Printf("O'zgartirilgan qiymatlar: a = %d, b = %d\n", a, b)
// }

func main() {
	var a, b int
	fmt.Println("Sonlarni kiriting:")
	fmt.Scan(&a)
	fmt.Scan(&b)

	if a != b  {
		a+=b
		fmt.Println(" qiymat: a = ", a)
	}else if a == b{
		c := 0
		fmt.Println(c)
	}
}