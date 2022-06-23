package request

type Register struct {
	Username  string `json:"userName" binding:"required,min=1,max=50"`
	Password  string `json:"password" binding:"required,min=6,max=16"`
	NickName  string `json:"nickName"`
	Mobile    string `json:"phone" binding:"required,min=11,max=11"`
	Email     string `json:"email"`
	HeaderImg string `json:"headerImg"`
}

type ChangeUserInfo struct {
	ID        int
	NickName  string `json:"nickName"`
	Mobile    string `json:"phone"`
	Email     string `json:"email"`
	HeaderImg string `json:"headerImg"`
}
