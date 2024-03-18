package common

type PermissionError struct{}

func (e PermissionError) Error() string {
	return "Permission denied"
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "Not found"
}

type ValidationError struct {
	message string
}

func (e ValidationError) Error() string {
	return e.message
}

func NewValidationError(message string) error {
	return ValidationError{
		message: message,
	}
}
