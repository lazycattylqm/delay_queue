package test

import (
	"com.lqm.demo/delay_queue"
	"testing"
)

func TestNew(t *testing.T) {
	dq := delay_queue.New[any]()
	if dq == nil {
		t.Error("New() should not be nil")
	}
}
