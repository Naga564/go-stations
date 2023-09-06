package model

type ErrNotFound struct {
	msg string
}

func (e *ErrNotFound) Error() string {
	return e.msg
}

func SetErrNotFound(message string) *ErrNotFound {
	return &ErrNotFound{msg: message}
}
