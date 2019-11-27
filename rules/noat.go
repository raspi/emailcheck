package rules

import "strings"

type ErrNoAt struct {
}

func (e *ErrNoAt) Error() string {
	return `empty`
}

func (e ErrNoAt) Validate(email string) (errs []error) {
	if strings.Count(email, `@`) != 1 {
		return append(errs, &e)
	}
	return errs
}
