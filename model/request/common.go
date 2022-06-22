package request

type PageInfo struct {
	Page     int    `json:"page" form:"page" binding:"required"`         //页码
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required"` //每页大小
	Keyword  string `json:"keyword" form:"keyword"`                      //关键字
}

type ById struct {
	ID int `json:"id" form:"id" binding:"required"` // 主键ID
}
