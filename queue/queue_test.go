package queue

import (
	"com.lqm.go.demo/item"
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	queue := New[string]()
	if len(queue.Items) != 0 {
		t.Errorf("Expected queue to be empty")
	}
	if queue.C == nil {
		t.Errorf("Expected queue to have a channel")
	}
}

func TestAddItem(t *testing.T) {
	queue := New[string]()
	itemA := item.Item[string]{
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
	queue := New[string]()
	itemA := item.Item[string]{
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
	queue := New[string]()
	itemA := item.Item[string]{
		Id:     "1",
		Expire: 3000,
		Data:   "test",
	}
	itemB := item.Item[string]{
		Id:     "2",
		Expire: 3000,
		Data:   "test",
	}
	itemC := item.Item[string]{
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

func TestUpdate(t *testing.T) {
	queue := New[string]()
	itemA := item.Item[string]{
		Id:     "1",
		Expire: 3000,
		Data:   "test",
	}
	itemB := item.Item[string]{
		Id:     "2",
		Expire: 3000,
		Data:   "test",
	}
	_ = queue.Add(itemA)
	_ = queue.Add(itemB)

	queue.UpdateItem(
		item.Item[string]{
			Id:     "1",
			Expire: 3000,
			Data:   "test",
		}, func(e1, e2 item.Item[string]) item.Item[string] {
			data := fmt.Sprint(e1.Data)
			data2 := fmt.Sprint(e2.Data)
			return item.Item[string]{
				Id:     e2.Id,
				Expire: e2.Expire + 1000,
				Data:   data + data2,
			}
		},
	)
	if queue.Items[0].Data != "testtest" {
		t.Errorf("expected as testtest")
	}
}

func TestQueue_Offer(t *testing.T) {
	itemA := item.Item[string]{
		Id:     "1",
		Expire: 3000,
		Data:   "test",
	}

	itemB := item.Item[string]{
		Id:     "2",
		Expire: 3000,
		Data:   "test2",
	}

	itemC := item.Item[string]{
		Id:     "1",
		Expire: 3000,
		Data:   "testc",
	}

	itemD := item.Item[string]{
		Id:     "3",
		Expire: 3000,
		Data:   "testd",
	}

	queue := New[string]()
	_ = queue.Add(itemA)
	_ = queue.Add(itemB)
	queue.Offer(
		itemC, func(e1, e2 item.Item[string]) item.Item[string] {
			data1 := fmt.Sprint(e1.Data)
			data2 := fmt.Sprint(e2.Data)
			return item.Item[string]{
				Id:     e1.Id,
				Expire: e1.Expire,
				Data:   data1 + " " + data2,
			}
		},
	)
	if queue.Items[0].Data != "test testc" {
		t.Errorf("error expected test test")
	}

	queue.Offer(
		itemD, func(e1, e2 item.Item[string]) item.Item[string] {
			data1 := fmt.Sprint(e1.Data)
			data2 := fmt.Sprint(e2.Data)
			return item.Item[string]{
				Id:     e1.Id,
				Expire: e1.Expire,
				Data:   data1 + " " + data2,
			}
		},
	)
	if len(queue.Items) != 3 {
		t.Errorf("expect len as 3 but now is %v", len(queue.Items))
	}

}

func TestQueue_Take(t *testing.T) {
	after := time.After(5 * time.Second)
	itemA := item.NewOne("1", 2000, "test")

	itemB := item.NewOne("2", 2000, "testb")

	itemC := item.NewOne("3", 1000, "testc")

	queue := New[string]()
	_ = queue.Add(itemA)
	_ = queue.Add(itemB)
	_ = queue.Add(itemC)
	queue.Take()

	var out item.Item[string]
	for {
		select {
		case out = <-queue.C:
			fmt.Printf("%v %v \n", time.Now(), out.Data)
		case <-after:
			fmt.Println("time out finish")
			return
		case <-queue.F:
			fmt.Println("finish for queue empty")
		}

	}
}
