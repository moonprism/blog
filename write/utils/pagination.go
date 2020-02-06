package utils

type Pagination struct {
	Page  int	`json:"page"`
	Limit int	`json:"limit"`
	Total int64	`json:"total"`
}

func GeneratePagination(page, pageSize int, count int64) (pagination *Pagination) {
	pagination = &Pagination{
		Page:  page,
		Limit: pageSize,
		Total: count,
	}
	return
}