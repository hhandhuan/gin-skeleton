package errors

const (
	AuthCode   = 4001
	ParamCode  = 4002
	ServerCode = 5000
	CommonCode = 4000
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
