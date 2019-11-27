package rules

import (
	"strings"
)

type ErrNoDotAfterAt struct {
}

func (e *ErrNoDotAfterAt) Error() string {
	return `no dot (.) after at (@)`
}

func (e ErrNoDotAfterAt) Validate(email string) (errs []error) {
	atIndex := strings.Index(email, `@`)
	dotIndex := strings.LastIndex(email, `.`)

	if atIndex > 0 && dotIndex > 0 && dotIndex > atIndex {
		return errs
	}

	return append(errs, &e)
}
