package delay_queue

import (
	"com.lqm.demo/queue"
	"com.lqm.go.demo/item"
	"context"
	"fmt"
	"sync"
	"time"
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

func (dq *DelayQueue[T]) GetQueue() *queue.Queue[T] {
	return dq.queue
}

func (dq *DelayQueue[T]) OfferTask(item item.Item[T], f func(old, new item.Item[T]) item.Item[T]) {
	dq.mu.Lock()
	defer dq.mu.Unlock()
	dq.queue.Offer(item, f)
}

func (dq *DelayQueue[T]) Run() {
	dq.queue.Take()
}

func (dq *DelayQueue[T]) ExeFuncWhenDone(after <-chan time.Time, f func(id string, data T), doneToStop bool) {
	ctx, cancelFunc := context.WithCancel(context.TODO())
	go func() {
		for {
			select {
			case outItem := <-dq.queue.C:
				f(outItem.Id, outItem.Data)
			case <-after:
				fmt.Println("for time out finish")
				if doneToStop {
					cancelFunc()
				}

			}
		}
	}()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("finished by user or other reason")
		}
	}()

}
