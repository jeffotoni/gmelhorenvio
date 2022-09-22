package models

type HttpErr struct {
	Error string `json:"error,omitempty"`
}

func NewHttpErr(err error) HttpErr {
	return HttpErr{Error: err.Error()}
}

func NewHttpErrStr(err string) HttpErr {
	return HttpErr{Error: err}
}
