

package main

import "fmt"

type Data struct {
	value int
	items []int
}

func modify(d Data){
	d.value = 10
	d.items[0] = 10

}
func main(){
	data := Data{
		value: 0,
		items: []int{0,1,2},
	}
	modify(data)

	fmt.Println(data.value)
	fmt.Println(data.items[0])
}