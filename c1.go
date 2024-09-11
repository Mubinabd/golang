package main

func f() byte {
	return 1 << int(9) >> int(9)
}
func g[Int int]() byte {
	return 1 << Int(9) >> Int(9)
}
func main() {
	println(f() == g())
}
