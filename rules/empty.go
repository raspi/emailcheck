package rules

type ErrEmpty struct {
}

func (e *ErrEmpty) Error() string {
	return `empty`
}

// Validate checks that given email address is not empty string
func (e ErrEmpty) Validate(email string) (errs []error) {
	if email == `` {
		return append(errs, &e)
	}

	return errs
}
