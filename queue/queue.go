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
}

func New[T any]() *Queue[T] {
	c := make(chan item.Item[T], 10)
	items := make([]*item.Item[T], 0)
	return &Queue[T]{
		C:     c,
		Items: items,
	}
}

func (q *Queue[T]) Find(item2 item.Item[T]) (result *item.Item[T]) {

	if len(q.Items) == 0 {
		return nil
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
	if q.Find(item2) != nil {
		return errors.New("item already exists")
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Items = append(q.Items, &item2)
	return nil
}

func (q *Queue[T]) DeleteItem(value item.Item[T]) {
	if q.Find(value) == nil {
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
	if q.Find(value) == nil {
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

func (q *Queue[T]) UpdateItem(item2 item.Item[T], f func(e1, e2 *item.Item[T]) *item.Item[T]) {
	find := q.Find(item2)
	if find == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	if f == nil {
		find = &item2
	} else {
		newOne := f(find, &item2)
		find.Data = newOne.Data
	}

}

func (q *Queue[T]) Offer(value item.Item[T], merge func(e1, e2 *item.Item[T]) *item.Item[T]) {
	if q.Find(value) == nil {
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
			for _, i := range q.Items {
				if i.Expired() {
					q.DeleteItem(*i)
					q.C <- *i
				}
			}
		}
	}()
}
