package response

import (
	"github.com/Dlimingliang/go-admin/model"
)

type UserResult struct {
	User model.User `json:"user"` //用户信息
}

type LoginResponse struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}
