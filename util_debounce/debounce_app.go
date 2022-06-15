package main

import (
	"com.lqm.demo/delay_queue"
	_struct "com.lqm.demo/util_debounce/struct"
	"com.lqm.go.demo/item"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	timeout, _ := context.WithTimeout(context.Background(), time.Second*10)

	node := _struct.DataNode[string]{
		Count: 1,
		Data:  "this is node 1",
	}

	newNode := _struct.DataNode[string]{
		Count: 1,
		Data:  "this is node 2",
	}

	i := item.New("1", 3000, node)
	i2 := item.New("2", 1000, newNode)
	fmt.Printf("before start time is %v\n", time.Now().String())
	d := delay_queue.New[_struct.DataNode[string]]()
	d.Run()
	d.OfferTask(
		*i, func(old, new _struct.DataNode[string]) _struct.DataNode[string] {
			return _struct.Merge(
				old, new, func(e1, e2 string) string {
					return e2
				},
			)
		},
	)
	d.OfferTask(
		*i2, func(old, new _struct.DataNode[string]) _struct.DataNode[string] {
			return _struct.Merge(
				old, new, func(e1, e2 string) string {
					return e2
				},
			)
		},
	)
	d.OfferTask(
		*i2, func(old, new _struct.DataNode[string]) _struct.DataNode[string] {
			return _struct.Merge(
				old, new, func(e1, e2 string) string {
					return e2
				},
			)
		},
	)
	d.OfferTask(
		*i2, func(old, new _struct.DataNode[string]) _struct.DataNode[string] {
			return _struct.Merge(
				old, new, func(e1, e2 string) string {
					return e2
				},
			)
		},
	)
	d.ExeFuncWhenDone(
		time.After(time.Second*10), func(id string, dataNode _struct.DataNode[string]) {
			marshal, err := json.Marshal(dataNode)
			var dataString string
			if err != nil {
				dataString = err.Error()
			} else {
				dataString = string(marshal)
			}
			fmt.Printf("time is %v, id is %v, data is %v\n", time.Now().String(), id, dataString)

		}, false,
	)
	select {
	case <-timeout.Done():
		fmt.Println("finish main")
		fmt.Println(time.Now().String())
	}
}
