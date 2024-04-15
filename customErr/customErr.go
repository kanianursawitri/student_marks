package customErr

type CustomError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

func (e CustomError) Error() string {
	return e.Message
}
func (e CustomError) Status() int {
	return e.StatusCode
}

func NewBadRequestError(message string) CustomError {
	return CustomError{Message: message, StatusCode: 400}
}

func NewNotFoundError(message string) CustomError {
	return CustomError{Message: message, StatusCode: 404}
}

func NewInternalServerError(message string) CustomError {
	return CustomError{Message: message, StatusCode: 500}
}
