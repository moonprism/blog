package se

type SearchEngine interface {
	Insert(index string, id int64, data interface{}) error
	Update(index string, id int64, data interface{}) error
	Delete(index string, id int64) error
	Search(index, text string, offset, limit int) (result []interface{}, sum int64, err error)
}

var Engine SearchEngine
