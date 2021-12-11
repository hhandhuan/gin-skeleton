package errors

const (
	UNAuthCode = 4001 // 未授权错误
	ParamCode  = 4002 // 参数错误
	ServerCode = 5000 // 服务端错误
	CommonCode = 4000 // 通用错误
)

type Error struct {
	Code int
	Msg  string
}

func NewError(code int, err interface{}) *Error {
	if v, ok := err.(error); ok {
		err = v.Error()
	}
	return &Error{Code: code, Msg: err.(string)}
}
