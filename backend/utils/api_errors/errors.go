package api_errors

type Error struct {
	Err        error
	statusCode int
	message    string
}

func New(err error, statusCode int, message string) error {
	return &Error{
		Err:        err,
		statusCode: statusCode,
		message:    message,
	}
}

func (e Error) Error() string {
	if e.message == "" {
		return e.Err.Error()
	}
	return e.message + " : " + e.Err.Error()
}

func (e Error) Unwrap() error {
	return e.Err
}

func (e Error) StatusCode() int {
	return e.statusCode
}
