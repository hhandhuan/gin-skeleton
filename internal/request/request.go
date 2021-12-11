package request

type (
	CreateAuthTokenRequest struct {
		Username string `json:"username" binding:"required"` // 登陆用户
		Password string `json:"password" binding:"required"` // 登陆密码
	}
	EditUserRequest struct {
		Username string `json:"username" binding:"required"` // 登录名
		Nickname string `json:"nickname" binding:"required"` // 用户别名
	}
)
