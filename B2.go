package main

import "fmt"

func main() {
	// for i := 0; i < 140; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// for i := 100; i < 240; i++ {
	// 	if i%2 !=0{
	// 		fmt.Println(i)
	// 	}
	// }

	// for i :=40;i<180;i++{
	// 	if i %7 == 0{
	// 		fmt.Println(i)
	// 	}
	// }

	// for i :=50;i<490;i++{
	// 	if i %11 == 0{
	// 		fmt.Println(i)
	// 	}
	// }

	// for i :=90;i<210;i++{
	// 	if i %25 == 0{
	// 		fmt.Println(i)
	// 	}
	// }

	// for i :=25;i<690;i++{
	// 	if i %5 == 0 || i %9 ==0{
	// 		fmt.Println(i)
	// 	}
	// }

	// for i :=1;i<4000;i++{
	// 	if i %11 == 0 || i %3 == 0 || i %9 == 0{
	// 		fmt.Println(i)
	// 	}
	// }

	// var k ,n int

	// fmt.Scan(&k)
	// fmt.Scan(&n)

	// if n > 0{

	// 	for i := 1; i <= n; i++ {
	// 		fmt.Println(k )
	// 	}
	// }

	// var a, b int
	// fmt.Scan(&a)
	// fmt.Scan(&b)

	// if a < b {
	// 	count := 0
	// 	for i := a; i <= b; i++ {
	// 		fmt.Println(i)
	// 		count++
	// 	}
	// 	fmt.Println("Sonlar soni:", count)
	// } else {
	// 	fmt.Println("a < b bo'lishi kerak")
	// }

	// var a, b int
	// fmt.Scan(&a)
	// fmt.Scan(&b)

	// if a < b {
	// 	count := 0
	// 	for i := b; i >= a; i-- {
	// 		fmt.Println(i)
	// 		count++
	// 	}
	// 	fmt.Println("Sonlar soni:", count)
	// } else {
	// 	fmt.Println("a < b bo'lishi kerak")
	// }


	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if n %i == 0{
			fmt.Println(i)
		}
	}
}
