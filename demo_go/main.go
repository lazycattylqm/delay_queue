package main

import (
	"com.lqm.go.demo/item"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	ctx, cancelFunc := context.WithCancel(context.TODO())

	go func() {
		for i := 0; i < 10; i++ {
			type Student struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
				Age  int    `json:"age"`
				Rate int    `json:"rate"`
			}
			student := Student{
				Id:   i,
				Name: "lqm",
				Age:  18,
				Rate: 100,
			}

			itemValue := item.Item[Student]{
				Id:     strconv.Itoa(i),
				Data:   student,
				Expire: 3000,
				Unit:   time.Millisecond,
			}

			marshal, err := json.Marshal(itemValue)
			if err != nil {
				cancelFunc()
			}
			fmt.Printf("Hello %v \n", string(marshal))

		}
		cancelFunc()
	}()
	fmt.Println("wait for the end")
	select {
	case <-ctx.Done():
		fmt.Println("cancel and done")
	}
}
