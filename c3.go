package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	resultChan := make(chan int, 1)

	go func() {
		result := simulateRequest()
		resultChan <- result
	}()

	select {
	case res := <-resultChan:
		fmt.Printf("received response: %d\n", res)
	case <-ctx.Done():
		panic("timeout occurred")
	}
}

func simulateRequest() int {
	time.Sleep(5 * time.Second)
	return 1
}
