package test_demo

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestName(t *testing.T) {
	mutex := sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	var flag int32 = 2
	swapInt32 := atomic.CompareAndSwapInt32(&flag, 0, 1)
	fmt.Println(swapInt32)
	fmt.Println(flag)

}

func TestName2(t *testing.T) {
	mutex := sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	var flag int32 = 2
	swapInt32 := atomic.CompareAndSwapInt32(&flag, 2, 1)
	fmt.Println(swapInt32)
	fmt.Println(flag)

}

func TestName3(t *testing.T) {
	mutex := sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	var flag int32 = 2
	old := atomic.SwapInt32(&flag, 1)
	fmt.Println(old)
	swapInt32 := atomic.CompareAndSwapInt32(&old, 1, 1)
	fmt.Println(swapInt32)
	fmt.Println(flag)

}
