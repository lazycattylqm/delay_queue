package test

import (
	"com.lqm.demo/delay_lock/delay_lock"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	mutex := delay_lock.New(time.Duration(3) * time.Second)
	if mutex == nil {
		t.Errorf("it should not be nil")
	}
}
