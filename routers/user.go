package routers

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

type UserRouter struct{}

func (userRouter *UserRouter) InitUserRouter(group *gin.RouterGroup) {
	userApi := v1.ApiGroupAPP.UserApi
	userGroup := group.Group("user")
	{
		userGroup.POST("admin_register", userApi.RegisterAdmin)          //注册管理员
		userGroup.PUT("setUserInfo", userApi.SetUserInfo)                //设置用户信息
		userGroup.POST("resetPassword", userApi.ResetPassword)           //重置密码为123456
		userGroup.DELETE("deleteUser", userApi.DeleteUser)               //删除用户
		userGroup.POST("setUserAuthorities", userApi.SetUserAuthorities) //设置用户角色
	}

	userGroupWithoutRecord := group.Group("user")
	{
		userGroupWithoutRecord.POST("getUserList", userApi.GetUserList) //分页获取用户列表
	}
}
