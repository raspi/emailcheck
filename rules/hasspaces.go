package rules

import "strings"

type ErrHasSpaces struct {
}

func (e *ErrHasSpaces) Error() string {
	return `spaces in address`
}

// Validate checks that there are no spaces in given email address
func (e ErrHasSpaces) Validate(email string) (errs []error) {
	if strings.Count(email, ` `) != 0 {
		return append(errs, &e)
	}

	return errs
}
