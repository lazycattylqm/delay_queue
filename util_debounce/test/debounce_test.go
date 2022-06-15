package test

import (
	_struct "com.lqm.demo/util_debounce/struct"
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	node := _struct.DataNode[string]{
		Data: "this is node 1",
	}

	node2 := _struct.DataNode[string]{
		Data: "this is node 2",
	}

	node3 := _struct.Merge(
		node, node2, func(e1, e2 string) string {
			return e1 + " " + e2
		},
	)
	fmt.Println(node3)
}
