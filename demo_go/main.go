package main

import (
	"com.lqm.go.demo/item"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type A struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (aa *A) String() string {
	marshal, err := json.Marshal(aa)
	if err != nil {
		return err.Error()
	}
	return string(marshal)
}

func (aa A) NameforA() {
	aa.Name = "lqiming"
}

func main() {

	TestB()

}

func TestB() {
	a := &A{
		Name: "lqm",
		Age:  18,
	}
	b := &A{
		Name: "fxl",
		Age:  17,
	}
	as := make([]*A, 0)
	as = append(as, a)
	as = append(as, b)
	fmt.Printf("as is %v \n", as)
	a.Name = "lqm2"
	fmt.Printf("as is %v \n", as)
}

func simpleTestA() {
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
