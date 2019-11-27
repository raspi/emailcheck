package rules

type ErrEmpty struct {
}

func (e *ErrEmpty) Error() string {
	return `empty`
}

func (e ErrEmpty) Validate(email string) (errs []error) {
	if email == `` {
		return append(errs, &e)
	}

	return errs
}
