package core

type VirenmonitoringError struct {
	IsVirenmonitoringError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewVirenmonitoringError(code string, msg string, ctx *Context) *VirenmonitoringError {
	return &VirenmonitoringError{
		IsVirenmonitoringError: true,
		Sdk:              "Virenmonitoring",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *VirenmonitoringError) Error() string {
	return e.Msg
}
