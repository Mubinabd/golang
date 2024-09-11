

package main

import "fmt"

func main(){
	data := make(chan int)
	shutdown := make (chan int)
	close(shutdown)
	close(data)

	select {
	case <- shutdown:
		fmt.Println("CLOSED, ")
	case data <- 1:
		fmt.Println("HAS WRITTEN, ")
	default:
		fmt.Println("DEFAULT, ")
	}
}