package core

// CreateSlice 创建指定长度并初始化为相同值的切片
func CreateSlice[T int | string | interface{}](length int, value T) []T {
	slice := make([]T, length)
	for i := range slice {
		slice[i] = value
	}
	return slice
}
