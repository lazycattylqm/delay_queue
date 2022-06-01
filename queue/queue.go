package queue

import (
	"com.lqm.go.demo/item"
	"errors"
	"fmt"
	"sync"
)

type Queue struct {
	C     chan item.Item[any]
	Items []*item.Item[any]
	mu    sync.RWMutex
}

func New() *Queue {
	c := make(chan item.Item[any], 10)
	items := make([]*item.Item[any], 0)
	return &Queue{
		C:     c,
		Items: items,
	}
}

func (q *Queue) Find(item2 item.Item[any]) (result *item.Item[any]) {

	if len(q.Items) == 0 {
		return nil
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	return FindByItem(
		q.Items, func(item item.Item[any]) bool {
			return item.EqualId(item2)
		},
	)
}

func (q *Queue) Add(item2 item.Item[any]) error {
	if q.Find(item2) != nil {
		return errors.New("item already exists")
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	q.Items = append(q.Items, &item2)
	return nil
}

func (q *Queue) DeleteItem(value item.Item[any]) {
	if q.Find(value) == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	byItem := RejectByItem(
		q.Items, func(a item.Item[any]) bool {
			return a.EqualId(value)
		},
	)
	q.Items = byItem
}

func (q *Queue) FilterItems(value item.Item[any]) {
	if q.Find(value) == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	byItem := FilterByItem(
		q.Items, func(a item.Item[any]) bool {
			return a.EqualId(value)
		},
	)
	q.Items = byItem
}

func (q *Queue) UpdateItem(item2 item.Item[any], f func(e1, e2 *item.Item[any]) *item.Item[any]) {
	find := q.Find(item2)
	if find == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	if f == nil {
		find = &item2
	} else {
		find = f(find, &item2)
	}

}

func (q *Queue) Offer(value item.Item[any], merge func(e1, e2 *item.Item[any]) *item.Item[any]) {
	if q.Find(value) == nil {
		err := q.Add(value)
		if err != nil {
			_ = fmt.Errorf("add item error: %s", err.Error())
		}
	}
	q.UpdateItem(value, merge)
}

func (q *Queue) Take() {
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
