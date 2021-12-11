package request

type (
	CreateAuthTokenRequest struct {
		Username string `json:"username" binding:"required"` // 登陆用户
		Password string `json:"password" binding:"required"` // 登陆密码
	}
)
