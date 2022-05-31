package main

import (
	"com.lqm.go.demo/item"
	"context"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	ctx, cancelFunc := context.WithCancel(context.TODO())
	itemValue := item.Item{
		Id: "123",
	}
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Hello, World! %v \n", itemValue)

		}
		cancelFunc()
	}()
	fmt.Println("wait for the end")
	select {
	case <-ctx.Done():
		fmt.Println("cancel and done")
	}
}
