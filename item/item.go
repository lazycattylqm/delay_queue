package item

import (
	"sync"
	"time"
)

type Item[T any] struct {
	Id     string `json:"id"`
	born   time.Time
	Data   T     `json:"data"`
	Expire int64 `json:"expire"`
	mu     sync.Mutex
	Unit   time.Duration `json:"unit"`
}

func New[T any](id string, expire int64, data T) *Item[T] {
	return &Item[T]{
		Id:     id,
		born:   time.Now(),
		Data:   data,
		Expire: expire,
		Unit:   time.Millisecond,
	}
}

func NewWithUnit[T any](id string, data T, expire int64, unit time.Duration) *Item[T] {
	return &Item[T]{
		Id:     id,
		born:   time.Now(),
		Data:   data,
		Expire: expire,
		Unit:   unit,
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
	escapeTime := calTimeDiff(*i)
	i.Expire = i.Expire - escapeTime
	if i.Expire < 0 {
		i.Expire = 0
	}
}

func (i *Item[T]) EqualId(another Item[T]) bool {
	return i.Id == another.Id
}

func (i *Item[T]) Expired() bool {
	i.mu.Lock()
	defer i.mu.Unlock()
	escapeTime := calTimeDiff(*i)
	i.Expire = i.Expire - escapeTime
	if i.Expire < 0 {
		i.Expire = 0
	}
	return i.Expire == 0
}

func calTimeDiff[T any](i Item[T]) int64 {
	timeDif := time.Now().Sub(i.born)
	escapeTime := timeDif / i.Unit
	return int64(escapeTime)
}
