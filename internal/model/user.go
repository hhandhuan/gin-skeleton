package model

type User struct {
	ID        int64  `json:"id"`         // 用户 ID
	Email     string `json:"phone"`      // 用户手机
	Password  string `json:"-"`          // 用户密码
	State     int    `json:"state"`      // 用户呢称
	CreatedAt *cTime `json:"created_at"` // 创建时间
	UpdatedAt *cTime `json:"updated_at"` // 更新时间
}
