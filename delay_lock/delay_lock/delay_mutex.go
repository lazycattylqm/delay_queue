package delay_lock

import (
	"context"
	"sync"
	"sync/atomic"
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

func (m *DelayMutex) Lock() {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

}

func (m *DelayMutex) checkForLock() bool {
	canLock := atomic.CompareAndSwapInt32(&m.flag, 0, 1)
	if !canLock {
		return false
	}
	m.flag = 1
	go func() {
		select {
		case <-m.timeOut.Done():
			m.channel <- struct{}{}
		}
	}()
	return true
}

func (m *DelayMutex) changeForUnlock() {
	canUnLock := atomic.CompareAndSwapInt32(&m.flag, 1, 0)
	if !canUnLock {
		return
	}
	m.flag = 0
	select {
	case <-m.channel:
		return
	}

}
