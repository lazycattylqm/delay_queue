package _struct

type DataNode[T any] struct {
	Count int `json:"count"`
	Data  T   `json:"data"`
}

func Merge[T any](old, new DataNode[T], f func(e1, e2 T) T) DataNode[T] {
	return DataNode[T]{
		Count: old.Count + 1,
		Data:  f(old.Data, new.Data),
	}
}
