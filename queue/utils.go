package queue

func FindByItem[T any](items []*T, f func(e1 T) bool) (int, *T) {
	for index, element := range items {
		if f(*element) {
			return index, element
		}
	}
	return -1, nil
}

func FilterByItem[T any](items []*T, f func(e1 T) bool) []*T {
	newItems := make([]*T, 0)
	for _, item := range items {
		if f(*item) {
			newItems = append(newItems, item)
		}
	}
	return newItems
}

func RejectByItem[T any](items []*T, f func(e T) bool) []*T {
	newItems := make([]*T, 0)
	for _, item := range items {
		if !f(*item) {
			newItems = append(newItems, item)
		}
	}
	return newItems
}
