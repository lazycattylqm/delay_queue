package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	ctx, cancelFunc := context.WithCancel(context.TODO())
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Hello, World! %v \n", i)

		}
		cancelFunc()
	}()
	fmt.Println("wait for the end")
	select {
	case <-ctx.Done():
		fmt.Println("cancel and done")
	}
}
