package main

import "fmt"

func main() {
	m := Meeseeks{}
	m.SetRequest("Helo there")
	fmt.Println("%s ",m)
}
type Meeseeks struct {
	request string
}

func (m Meeseeks)SetRequest(request string){
	m.request = request
}
func (m Meeseeks)String()string{
	return m.request
}
