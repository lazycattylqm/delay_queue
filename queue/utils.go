package queue

func FindByItem[T any](items []*T, item2 T, f func(e1, e2 T) bool) *T {
	for _, i := range items {
		if f(*i, item2) {
			return i
		}
	}
	return nil
}

func FilterByItem[T any](items []*T, target T, f func(e1, e2 T) bool) []*T {
	newItems := make([]*T, len(items))
	for _, item := range items {
		if f(*item, target) {
			newItems = append(newItems, item)
		}
	}
	return newItems
}
