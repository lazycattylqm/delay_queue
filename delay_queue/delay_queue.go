package delay_queue

import (
	"com.lqm.demo/queue"
	"com.lqm.go.demo/item"
	"sync"
)

type DelayQueue[T any] struct {
	queue *queue.Queue[T]
	mu    sync.RWMutex
}

func New[T any]() *DelayQueue[T] {
	q := queue.New[T]()
	return &DelayQueue[T]{
		queue: q,
	}
}

func (dq *DelayQueue[T]) OfferTask(item item.Item[T], f func(old, new item.Item[T]) item.Item[T]) {
	dq.mu.Lock()
	defer dq.mu.Unlock()
	dq.queue.Offer(item, f)
}
