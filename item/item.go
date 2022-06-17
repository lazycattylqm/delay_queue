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

func NewOne[T any](id string, expire int64, data T) Item[T] {
	return Item[T]{
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

func (i *Item[T]) UpdateWithFunc(data T, updateFunc func(old, new T) T) {
	//defer i.Expired()
	i.updateData(data, updateFunc)
}

func (i *Item[T]) Update(data T) {
	defer i.Expired()
	i.updateData(data, nil)
}

func (i *Item[T]) updateData(data T, updateFunc func(old, new T) T) {
	//i.mu.Lock()
	//defer i.mu.Unlock()
	if i.checkExpired() {
		return
	}
	if updateFunc == nil {
		i.Data = data
	} else {
		i.Data = updateFunc(i.Data, data)
	}

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
	if i.checkExpired() {
		i.Expire = 0
		return true
	}
	return false
}

func (i *Item[T]) checkExpired() bool {
	diff := calTimeDiff(*i)
	return i.Expire-diff <= 0
}

func calTimeDiff[T any](i Item[T]) int64 {
	timeDif := time.Now().Sub(i.born)
	if i.Unit == 0 {
		// 如果是0 则立即返回expire设置时间 保证一定超时
		return i.Expire
	}
	escapeTime := timeDif / i.Unit
	return int64(escapeTime)
}
