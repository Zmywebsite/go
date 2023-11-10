package user

type UserSchema struct {
	UserName           string `json:"userName"`           // 用户名
	Password           string `json:"password"`           // 密码
	CreatTime          int    `json:"creatTime"`          // 创建时间
	UpdateTime         int    `json:"updateTime"`         // 更新时间
	LastLoginTime      int    `json:"lastLoginTime"`      // 最后一次登录时间
	FirstLogin         bool   `json:"firstLogin"`         // 首次登录
	PasswordUpdateTime int    `json:"passwordUpdateTime"` // 密码更新时间
}
