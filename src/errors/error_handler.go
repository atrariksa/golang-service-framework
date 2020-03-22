package errors

type Err interface {
	Error() string
}

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func New(text string) Err {
	return &errorString{text}
}
