package delay_lock

import (
	"context"
	"sync"
	"time"
)

type DelayMutex struct {
	sync.Mutex
	flag    int32
	channel chan struct{}
	timeOut context.Context
	cancelF context.CancelFunc
}

func New(duration time.Duration) *DelayMutex {
	timeout, cancelFunc := context.WithTimeout(context.TODO(), duration)
	c := make(chan struct{}, 1)
	return &DelayMutex{
		flag:    0,
		channel: c,
		timeOut: timeout,
		cancelF: cancelFunc,
	}
}
