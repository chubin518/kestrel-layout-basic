package result

type PageData[T any] struct {
	Total int `json:"total"`
	Index int `json:"index"`
	Size  int `json:"size"`
	List  []T `json:"list"`
}

// OfPage new a PageData
func OfPage[T any](list []T, total, index, size int) *PageData[T] {
	return &PageData[T]{
		List:  list,
		Total: total,
		Index: index,
		Size:  size,
	}
}
