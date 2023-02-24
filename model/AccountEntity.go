package model

type AccountEntity struct {
	BaseEntity
	Username string `gorm:"username" json:"username"` // 用户名
	Password string `gorm:"password" json:"password"` // 密码
	Role     string `gorm:"role" json:"role"`         // 身份
	Countid  string `gorm:"countid" json:"countid"`
	Coachid  string `gorm:"coachid" json:"coachid"`
}

type CoachEntity struct {
	Coachmame string `json:"mame"`
	Coachid   string `json:"id"`
}

func (t *AccountEntity) TableName() string {
	return "user"
}

type Cookie struct {
	Id       int    `gorm:"id" json:"id"`
	Username string `gorm:"username" json:"name"` // 用户名
	Token    string `gorm:"token" json:"token"`
	Role     string `gorm:"role" json:"role"` // 身份
}

func (t *Cookie) TableName() string {
	return "cookie"
}

type Sessions struct {
	Session  string `gorm:"session" json:"session"`
	Username string `gorm:"username" json:"username"` // 用户名
	Role     string `gorm:"role" json:"role"`         // 身份
}

func (t *Sessions) TableName() string {
	return "user"
}
