package queue

import "testing"

type tempS struct {
	Id int
}

func TestFindByItem(t *testing.T) {

	array := initData()
	_, item := FindByItem[tempS](
		array, func(e1 tempS) bool {
			return e1.Id == 5
		},
	)
	if item.Id != 5 {
		t.Errorf("Expected 5, got %d", *item)
	}

}

func TestFilterByItem(t *testing.T) {
	data := initData()
	item := FilterByItem(
		data, func(e1 tempS) bool {
			return e1.Id >= 5
		},
	)
	if item[0].Id < 5 {
		t.Errorf("Expected larger than 5, got %d", item[0].Id)
	}
}

func TestRejectByItem(t *testing.T) {
	data := initData()
	item := RejectByItem(
		data, func(e1 tempS) bool {
			return e1.Id >= 5
		},
	)
	if item[len(item)-1].Id >= 5 {
		t.Errorf("Expected smaller than 5, got %d", item[len(item)-1].Id)
	}
}

func initData() []*tempS {
	array := []*tempS{}
	for i := 0; i < 10; i++ {
		array = append(array, &tempS{Id: i})
	}
	return array
}
