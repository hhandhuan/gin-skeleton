package errors

const (
	AuthCode   = 4001
	ParamCode  = 4002
	ServerCode = 5000
	CommonCode = 4022
)

type Error struct {
	Code int
	Msg  string
}

func NewError(code int, err string) *Error {
	return &Error{Code: code, Msg: err}
}
