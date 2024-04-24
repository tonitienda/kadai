package common

type PermissionError struct{}

func (e PermissionError) Error() string {
	return "Permission denied"
}

type NotFoundError struct {
	message string
}

func (e NotFoundError) Error() string {
	return e.message
}

type ValidationError struct {
	message string
}

func (e ValidationError) Error() string {
	return e.message
}

type ForbiddenError struct {
	message string
}

func (e ForbiddenError) Error() string {
	return e.message
}

type StatusIncorrectError struct {
	message string
}

func (e StatusIncorrectError) Error() string {
	return e.message
}

func NewStatusIncorrectError(message string) error {
	return StatusIncorrectError{
		message: message,
	}
}

func NewValidationError(message string) error {
	return ValidationError{
		message: message,
	}
}

func NewForbiddenError(message string) error {
	return ForbiddenError{
		message: message,
	}
}

func NewNotFoundError(message string) error {
	return NotFoundError{
		message: message,
	}
}
