package error

type HttpError struct {
	ReturnCode string
	MessageKey string
	Message    string
}

func NewHttpError(returnCode, messageKey, message string) *HttpError {
	return &HttpError{
		ReturnCode: returnCode,
		MessageKey: messageKey,
		Message:    message,
	}
}

func (e *HttpError) Error() string {
	return e.Message
}
