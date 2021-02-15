package rules

import (
	"strings"
)

type ErrNoUser struct {
}

func (e *ErrNoUser) Error() string {
	return `no user`
}

// Validate checks that given email address has at least one character before @-sign
func (e ErrNoUser) Validate(email string) (errs []error) {
	atIndex := strings.Index(email, `@`)

	if atIndex <= 1 {
		return append(errs, &e)
	}

	return errs
}
