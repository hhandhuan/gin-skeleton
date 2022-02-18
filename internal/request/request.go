package request

type (
	CreateAuthTokenRequest struct {
		Email    string `json:"email" form:"username" binding:"required"`    // 登陆用户
		Password string `json:"password" form:"password" binding:"required"` // 登陆密码
	}
	EditUserRequest struct {
		Email    string `json:"email" form:"username" binding:"required"`    // 登录名
		Nickname string `json:"nickname" form:"nickname" binding:"required"` // 用户别名
	}
)
