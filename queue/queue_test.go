package queue

import (
	"com.lqm.go.demo/item"
	"testing"
)

func TestNew(t *testing.T) {
	queue := New()
	if len(queue.Items) != 0 {
		t.Errorf("Expected queue to be empty")
	}
	if queue.C == nil {
		t.Errorf("Expected queue to have a channel")
	}
}

func TestAddItem(t *testing.T) {
	queue := New()
	itemA := item.Item[any]{
		Id:     "1",
		Expire: 3000,
		Data:   "test",
	}
	if err := queue.Add(itemA); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(queue.Items) != 1 {
		t.Errorf("Expected queue to have 1 item")
	}

}

func TestDeleteCaseOne(t *testing.T) {
	queue := New()
	itemA := item.Item[any]{
		Id:     "1",
		Expire: 3000,
		Data:   "test",
	}
	if err := queue.Add(itemA); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	queue.DeleteItem(itemA)
	if len(queue.Items) != 0 {
		t.Errorf("Expected queue to be empty")
	}
}

func TestFilter(t *testing.T) {
	queue := New()
	itemA := item.Item[any]{
		Id:     "1",
		Expire: 3000,
		Data:   "test",
	}
	itemB := item.Item[any]{
		Id:     "2",
		Expire: 3000,
		Data:   "test",
	}
	itemC := item.Item[any]{
		Id:     "3",
		Expire: 3000,
		Data:   "test",
	}
	_ = queue.Add(itemA)
	_ = queue.Add(itemB)
	_ = queue.Add(itemC)
	queue.FilterItems(itemA)
	if len(queue.Items) != 1 {
		t.Errorf("Expected queue to have 1 items")
	}
}
