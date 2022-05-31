package item

import (
	"sync"
	"time"
)

type Item[T any] struct {
	Id     string
	Born   time.Time
	Data   T
	Expire int64
	mu     sync.Mutex
	unit   time.Duration
}

func New[T any](id string, expire int64, data T) *Item[T] {
	return &Item[T]{
		Id:     id,
		Born:   time.Now(),
		Data:   data,
		Expire: expire,
		unit:   time.Millisecond,
	}
}

func NewWithUnit[T any](id string, data T, expire int64, unit time.Duration) *Item[T] {
	return &Item[T]{
		Id:     id,
		Born:   time.Now(),
		Data:   data,
		Expire: expire,
		unit:   unit,
	}
}

func (i *Item[T]) UpdateWithFunc(data T, updateF func(old, new T) T) {
	i.updateDataAndExpireTime(data, updateF)
}

func (i *Item[T]) Update(data T) {
	i.updateDataAndExpireTime(data, nil)
}

func (i *Item[T]) updateDataAndExpireTime(data T, updateF func(old, new T) T) {
	i.mu.Lock()
	defer i.mu.Unlock()
	if updateF == nil {
		i.Data = data
	} else {
		i.Data = updateF(i.Data, data)
	}
	i.updateExpireTime()
}

func (i *Item[T]) updateExpireTime() {
	timeDif := time.Now().Sub(i.Born)
	escapeTime := timeDif / i.unit
	i.Expire = i.Expire - int64(escapeTime)
	if i.Expire < 0 {
		i.Expire = 0
	}
}

func (i *Item[T]) EqualId(another Item[T]) bool {
	return i.Id == another.Id
}
