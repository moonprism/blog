package core

import "strconv"

// CreateSlice 创建指定长度并初始化为相同值的切片
func CreateSlice[T int | string | interface{}](length int, value T) []T {
	slice := make([]T, length)
	for i := range slice {
		slice[i] = value
	}
	return slice
}

// ItoaSlice int 切片转换 string 切片
func ItoaSlice(originSlice []int) []string {
	strSlice := make([]string, len(originSlice))
	for i, v := range originSlice {
		strSlice[i] = strconv.Itoa(v)
	}
	return strSlice
}

// AtoISlice string 切片转换 int 切片
func AtoISlice(originSlice []string) (strSlice []int, err error) {
	strSlice = make([]int, len(originSlice))
	for i, v := range originSlice {
		strSlice[i], err = strconv.Atoi(v)
		if err != nil {
			return
		}
	}
	return
}
