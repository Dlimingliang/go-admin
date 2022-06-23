package business

func New(text string) error {
	return GAValidateError{text}
}

type GAValidateError struct {
	message string
}

func (ve GAValidateError) Error() string {
	return ve.message
}
