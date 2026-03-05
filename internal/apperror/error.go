package apperror

type AppError struct {
	Code       string
	Message    string
	StatusCode int
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code string, message string, status int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: status,
	}
}
