package item

import (
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
