package errorx

import (
	"errors"
	"fmt"
)

type Error struct {
	Code int64
	Msg  string
}

func (e Error) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.Code, e.Msg)
}

func NewError(code int64, msg string) Error {
	return Error{code, msg}
}

func (e Error) WithMessage(msg string) Error {
	e.Msg = msg
	return e
}

// ConvertErr convert error to Errno
func ConvertErr(err error) Error {
	Err := Error{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.Msg = err.Error()
	return s
}
