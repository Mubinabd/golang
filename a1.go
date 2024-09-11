package main

var true = false

const byte = 123

type nil interface{}

func len(nil) int {
	return byte
}

func main() {
	var s = []bool{true, true, true}
	println(s[0])
	println(len(s))
}
