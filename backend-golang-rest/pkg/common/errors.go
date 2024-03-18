package common

type PermissionError struct{}

func (e PermissionError) Error() string {
	return "Permission denied"
}

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "Not found"
}

type ValidationError struct{}

func (e ValidationError) Error() string {
	return "Validation error"
}
