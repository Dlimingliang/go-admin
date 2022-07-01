package request

type Register struct {
	Username     string   `json:"userName" binding:"required,min=1,max=50"` //用户名
	Password     string   `json:"password" binding:"required,min=6,max=16"` //密码
	NickName     string   `json:"nickName"`                                 //昵称
	Mobile       string   `json:"phone" binding:"required,min=11,max=11"`   //电话
	Email        string   `json:"email"`                                    //邮件
	HeaderImg    string   `json:"headerImg"`                                //头像
	AuthorityIds []string `json:"authorityIds"`                             //角色ID
}

type ChangeUserInfo struct {
	ID           int      //ID
	NickName     string   `json:"nickName"`     //昵称
	Mobile       string   `json:"phone"`        //电话
	Email        string   `json:"email"`        //邮件
	HeaderImg    string   `json:"headerImg"`    //头像
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}

type SetUserAuthorities struct {
	ID           int      `json:"ID"`
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}
