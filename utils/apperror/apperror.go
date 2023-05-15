package apperror

type AppError struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Errors interface{} `json:"errors"`
}

func (e *AppError) Error() string {
	return e.Msg
}
