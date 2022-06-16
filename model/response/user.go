package response

import (
	"time"
)

type UserListRes struct {
	PageRes
	List []UserRes `json:"list"`
}

type UserRes struct {
	ID         int       `json:"ID,omitempty"`
	CreateTime time.Time `json:"createTime"`
	CreateUser int       `json:"createUser,omitempty"`
	UpdateTime time.Time `json:"updateTime"`
	UpdateUser int       `json:"updateUser,omitempty"`
	Username   string    `json:"username,omitempty"`
	Password   string    `json:"password,omitempty"`
	NickName   string    `json:"nickName,omitempty"`
	Mobile     string    `json:"mobile,omitempty"`
	Email      string    `json:"email,omitempty"`
	HeaderImg  string    `json:"headerImg,omitempty"`
}
