package model

type User struct {
	ID        int64  `json:"id"`         // 用户 ID
	Phone     string `json:"phone"`      // 用户手机
	Password  string `json:"-"`          // 用户密码
	Username  string `json:"username"`   // 用户名称
	Nickname  string `json:"nickname"`   // 用户呢称
	CreateAt  int64  `json:"create_at"`  // 创建时间
	UpdatedAt int64  `json:"updated_at"` // 更新时间
}
