package model

type User struct {
	BaseModel
	Username  string `gorm:"column:username;type:varchar(50);not null;unique;comment:用户名"`
	Password  string `gorm:"column:password;type:varchar(100);not null;comment:密码"`
	NickName  string `gorm:"column:nick_name;type:varchar(100);default:'';comment:昵称"`
	Mobile    string `gorm:"column:mobile;type:varchar(11);not null;unique;comment:电话"`
	Email     string `gorm:"column:email;type:varchar(50);default:'';comment:邮箱"`
	HeaderImg string `gorm:"column:header_img;type:varchar(100);default:'';comment:头像"`
}
