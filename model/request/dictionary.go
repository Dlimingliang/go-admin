package request

import "github.com/Dlimingliang/go-admin/model"

type DictionarySearch struct {
	PageInfo
	model.Dictionary
}
