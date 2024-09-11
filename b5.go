package main

func main(){
	const (
		A,B = iota,iota
		C,_
		iota,_
		_,D
	)
	println(A,B,C,D)
}