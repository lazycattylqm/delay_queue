package queue

import (
	"com.lqm.go.demo/item"
	"errors"
	"fmt"
	"sync"
)

type Queue[T any] struct {
	C     chan item.Item[T]
	Items []*item.Item[T]
	mu    sync.RWMutex
	F     chan int
}

func New[T any]() *Queue[T] {
	c := make(chan item.Item[T], 10)
	f := make(chan int)
	items := make([]*item.Item[T], 0)
	return &Queue[T]{
		C:     c,
		Items: items,
		F:     f,
	}
}

func (q *Queue[T]) Find(item2 item.Item[T]) (index int, result *item.Item[T]) {

	if len(q.Items) == 0 {
		return -1, nil
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	return FindByItem(
		q.Items, func(item item.Item[T]) bool {
			return item.EqualId(item2)
		},
	)
}

func (q *Queue[T]) Add(item2 item.Item[T]) error {
	if _, ele := q.Find(item2); ele != nil {
		return errors.New("item already exists")
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Items = append(q.Items, &item2)
	return nil
}

func (q *Queue[T]) DeleteItem(value item.Item[T]) {
	if _, ele := q.Find(value); ele == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	byItem := RejectByItem(
		q.Items, func(a item.Item[T]) bool {
			return a.EqualId(value)
		},
	)
	q.Items = byItem
}

func (q *Queue[T]) FilterItems(value item.Item[T]) {
	if _, ele := q.Find(value); ele == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	byItem := FilterByItem(
		q.Items, func(a item.Item[T]) bool {
			return a.EqualId(value)
		},
	)
	q.Items = byItem
}

func (q *Queue[T]) UpdateItem(item2 item.Item[T], f func(e1, e2 item.Item[T]) item.Item[T]) {
	index, find := q.Find(item2)
	if find == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	newOne := f(*find, item2)
	q.Items[index] = &newOne

}

func (q *Queue[T]) Offer(value item.Item[T], merge func(e1, e2 item.Item[T]) item.Item[T]) {
	if _, ele := q.Find(value); ele == nil {
		err := q.Add(value)
		if err != nil {
			_ = fmt.Errorf("add item error: %s", err.Error())
		}
		return
	}
	q.UpdateItem(value, merge)
}

func (q *Queue[T]) Take() {
	go func() {
		for {
			if len(q.Items) == 0 {
				q.F <- -1
				return
			}
			for _, i := range q.Items {
				if i.Expired() {
					q.DeleteItem(*i)
					q.C <- *i
				}
			}
		}
	}()
}
