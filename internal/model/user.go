package model

import "github.com/hhandhuan/gin-skeleton/db"

var (
	UserTableName = "users"      // 表名
	UserModel     = &userModel{} // 模型
)

type (
	userModel  struct{}
	UserEntity struct {
		ID        int64  `json:"id"`
		Phone     string `json:"phone"`
		Password  string `json:"password"`
		Username  string `json:"username"`
		Nickname  string `json:"nickname"`
		CreateAt  int64  `json:"create_at"`
		UpdatedAt int64  `json:"updated_at"`
	}
)

func (u *userModel) GetUserByUsername(name string) *UserEntity {
	var user *UserEntity
	db.Mysql.Table(UserTableName).Where("username = ?", name).First(user)
	return user
}