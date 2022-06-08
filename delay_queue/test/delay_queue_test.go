package test

import (
	"com.lqm.demo/delay_queue"
	"com.lqm.go.demo/item"
	"testing"
)

func TestNew(t *testing.T) {
	dq := delay_queue.New[any]()
	if dq == nil {
		t.Error("New() should not be nil")
	}
}

func TestGetQueue(t *testing.T) {
	dq := delay_queue.New[any]()
	q := dq.GetQueue()
	if q == nil {
		t.Error("GetQueue() should not be nil")
	}
}

func TestOffer(t *testing.T) {
	dq := delay_queue.New[string]()
	i := item.New("id", 3000, "data")
	dq.OfferTask(
		*i, func(old, new item.Item[string]) item.Item[string] {
			return new
		},
	)
	if len(dq.GetQueue().Items) == 0 {
		t.Error("OfferTask() should not be nil")
	}

}

func TestOffer_Add_2_items(t *testing.T) {
	dq := delay_queue.New[string]()
	i := item.New("id", 3000, "data")
	dq.OfferTask(
		*i, func(old, new item.Item[string]) item.Item[string] {
			return new
		},
	)
	i2 := item.New("id2", 3000, "data")
	dq.OfferTask(
		*i2, func(old, new item.Item[string]) item.Item[string] {
			return new
		},
	)
	if len(dq.GetQueue().Items) != 2 {
		t.Error("OfferTask() should not be nil")
	}
}
