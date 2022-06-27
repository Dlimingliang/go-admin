package response

type PageResult struct {
	List     interface{} `json:"list"`     //列表
	Total    int64       `json:"total"`    //总数
	Page     int         `json:"page"`     //页码
	PageSize int         `json:"pageSize"` //每页大小
}
