package request

type (
	CreateAuthTokenRequest struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
	}
)
