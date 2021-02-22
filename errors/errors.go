package errors

//ServiceError to return error
type ServiceError struct {
	Code         int
	ErrorMessage string
}

func (ce *ServiceError) Error() string {
	return ce.ErrorMessage
}
