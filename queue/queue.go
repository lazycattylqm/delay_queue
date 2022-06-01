package queue

import (
	"com.lqm.go.demo/item"
	"errors"
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
		q.Items, item2, func(a, b item.Item[any]) bool {
			return a.EqualId(b)
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
	FilterByItem(
		q.Items, value, func(a, b item.Item[any]) bool {
			return a.EqualId(b)
		},
	)
}

func (q *Queue) UpdateItem(item2 item.Item[any], f func(e1, e2 *item.Item[any]) *item.Item[any]) {
	find := q.Find(item2)
	if find == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	find = f(find, &item2)
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
