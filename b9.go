package main

import (
	"cmp"
	"fmt"
)

func inCache(uid string) bool {
	fmt.Print("inCache")
	return true
}

func inDB(uid string) bool {
	fmt.Print("inDB")
	return true
}
func main() {
	uid := "elliot"
	cmp.Or(inCache(uid), inDB(uid))
}
