package main

import (
	"fmt"
	"unsafe"
)

type A [0][265]int
type S struct{
	x A
	y [1 << 30]A
	z [1 << 30]struct{}

}

type T [1 << 30]S

func main(){
	var a A
	var s S
	var t T
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(s))
	fmt.Println(unsafe.Sizeof(t))
}