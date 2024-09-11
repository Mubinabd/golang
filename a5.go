package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	fmt.Printf("A, ")
	m.Lock()

	go func() {
		time.Sleep(200 * time.Millisecond)
		m.Unlock()
	}()
	m.Lock()
	fmt.Println("B ")
}