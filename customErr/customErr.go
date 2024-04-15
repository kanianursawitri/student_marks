package customErr

type CustomError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (e CustomError) Error() string {
	return e.Msg
}
func (e CustomError) Status() int {
	return e.Code
}

func NewBadRequestError(message string) CustomError {
	return CustomError{Msg: message, Code: 400}
}

func NewNotFoundError(message string) CustomError {
	return CustomError{Msg: message, Code: 404}
}

func NewInternalServerError(message string) CustomError {
	return CustomError{Msg: message, Code: 500}
}
