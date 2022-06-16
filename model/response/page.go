package response

import "github.com/Dlimingliang/go-admin/model/common"

type PageRes struct {
	common.PageInfo
	Total int `json:"total"`
}
