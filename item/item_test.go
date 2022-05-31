package item

import (
	"encoding/json"
	"sync"
	"testing"
	"time"
)

func TestItemNew(t *testing.T) {
	i := New("1", 1000, "test")
	if i.Id != "1" {
		t.Errorf("Item ID should be 1, got %s", i.Id)
	}
	if i.Expire != 1000 {
		t.Errorf("Item price should be 1000, got %d", i.Expire)
	}
	if i.Data != "test" {
		t.Errorf("Item data should be test, got %s", i.Data)
	}
	t.Log("success")
}

func TestUpdate(t *testing.T) {
	i := New("1", 1000, "test")
	if i.Data != "test" {
		t.Errorf("Item data should be test, got %s", i.Data)
	}
	i.Update("testNew")
	if i.Data != "testNew" {
		t.Errorf("Item data should be test, got %s", i.Data)
	}
	t.Log("Success")
}

func TestUpdateWithFunc(t *testing.T) {
	i := New("1", 1000, "test")
	if i.Data != "test" {
		t.Errorf("Item data should be test, got %s", i.Data)
	}
	i.UpdateWithFunc(
		"test New", func(old, new string) string {
			return old + " " + new
		},
	)
	if i.Data != "test test New" {
		t.Errorf("Item data should be test, got %s", i.Data)
	}
}

func TestNotExpire(t *testing.T) {
	i := New("1", 3000, "test")
	time.Sleep(time.Duration(2) * time.Second)
	i.Update("test")
	if i.Expire > 1000 {
		t.Errorf("Item expire should be 1000, got %d", i.Expire)
	}

}

func TestExpire(t *testing.T) {
	i := New("1", 1000, "test")
	time.Sleep(time.Duration(2) * time.Second)
	i.Update("test")
	if i.Expire != 0 {
		t.Errorf("Item expire should be 1000, got %d", i.Expire)
	}
}

func TestItemEqualId(t *testing.T) {
	i := New("1", 1000, "test")
	i2 := New("1", 1000, "test")
	if !i.EqualId(*i2) {
		t.Errorf("Item should be equal, got %s", i.Id)
	}
}

func TestItemEqualIdNot(t *testing.T) {
	i := New("1", 1000, "test")
	i2 := New("2", 1000, "test")
	if i.EqualId(*i2) {
		t.Errorf("Item should be equal, got %s", i.Id)
	}
}

func TestNewWithUnitAndExpire(t *testing.T) {
	i := NewWithUnit("1", "test", 3, time.Second)
	i.Update("test2")
	if i.Data != "test2" {
		t.Errorf("Item data should be test2, got %s", i.Data)
	}
	time.Sleep(1 * time.Second)
	if i.Expired() {
		t.Errorf("Item should not be expired, got %s", i.Id)
	}
	if i.Expire > 2 {
		t.Errorf("Item expire should be less than 2, got %d", i.Expire)
	}
	time.Sleep(3 * time.Second)
	if !i.Expired() {
		t.Errorf("Item should be expired, got %s", i.Id)
	}
	if i.Expire != 0 {
		t.Errorf("Item expire should be 0, got %d", i.Expire)
	}
}

func TestMultiGoRoutine(t *testing.T) {
	i := New("1", 1000, "test")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		i.Update("test2")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(3) * time.Second)
		i.Update("test3")
	}()
	wg.Wait()
	marshal, err := json.Marshal(i)
	if err != nil {
		t.Errorf("Marshal error: %s", err)
	}
	t.Logf("%s", string(marshal))

}
