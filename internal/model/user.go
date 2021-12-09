package model

type User struct {
	ID        int64  `json:"id"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	CreateAt  int64  `json:"create_at"`
	UpdatedAt int64  `json:"updated_at"`
}
