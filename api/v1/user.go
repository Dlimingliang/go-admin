package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/common"
	"github.com/Dlimingliang/go-admin/model/response"
	"github.com/Dlimingliang/go-admin/service"
)

func UserToUserRes(user model.User) response.UserRes {
	return response.UserRes{
		ID:         user.ID,
		CreateTime: user.CreateTime,
		CreateUser: user.CreateUser,
		UpdateTime: user.UpdateTime,
		UpdateUser: user.UpdateUser,
		Username:   user.Username,
		Password:   user.Password,
		NickName:   user.NickName,
		Mobile:     user.Mobile,
		Email:      user.Email,
		HeaderImg:  user.HeaderImg,
	}
}

func GetUserList(ctx *gin.Context) {
	pageInfo := common.PageInfo{}
	//暂时不做错误处理
	_ = ctx.ShouldBind(&pageInfo)

	userList, total, _ := service.GetUserList(pageInfo)
	userRes := make([]response.UserRes, 0)
	for _, user := range userList {
		userRes = append(userRes, UserToUserRes(user))
	}

	ctx.JSON(http.StatusOK, response.Result{
		Code: 0,
		Msg:  "获取成功",
		Data: response.UserListRes{
			PageRes: response.PageRes{
				PageInfo: pageInfo,
				Total:    total,
			},
			List: userRes,
		},
	})
}
