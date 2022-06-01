package queue

import "testing"

func TestFindByItem(t *testing.T) {
	type tempS struct {
		Id int
	}
	array := []*tempS{}
	for i := 0; i < 10; i++ {
		array = append(array, &tempS{Id: i})
	}
	item := FindByItem[tempS](
		array, func(e1 tempS) bool {
			return e1.Id == 5
		},
	)
	if item.Id != 5 {
		t.Errorf("Expected 5, got %d", *item)
	}

}
