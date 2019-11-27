package rules

import (
	"strings"
)

type ErrNoUser struct {
}

func (e *ErrNoUser) Error() string {
	return `no user`
}

func (e ErrNoUser) Validate(email string) (errs []error) {
	atIndex := strings.Index(email, `@`)

	if atIndex <= 1 {
		return append(errs, &e)
	}

	return errs
}
