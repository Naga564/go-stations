package model

type ErrNotFound struct {
	msg string
}

func (e *ErrNotFound) Error() string {
	return e.msg
}
