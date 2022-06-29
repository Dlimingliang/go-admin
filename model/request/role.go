package request

type RoleInfo struct {
	AuthorityId   string `json:"authorityId" binding:"required,"`
	AuthorityName string `json:"authorityName" binding:"required,"`
	ParentId      string `json:"parentId" binding:"required,"`
}
