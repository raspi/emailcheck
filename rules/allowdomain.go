package rules

import (
	"fmt"
	"strings"
)

type ErrDomainNotInAllowedList struct {
	Domain         string
	Err            error
	allowedDomains []string
}

func NewErrDomainNotInAllowedList(allowed []string) ErrDomainNotInAllowedList {
	if len(allowed) == 0 {
		panic(`empty list`)
	}

	return ErrDomainNotInAllowedList{
		allowedDomains: allowed,
	}
}

func (e *ErrDomainNotInAllowedList) Error() string {
	return fmt.Sprintf(`domain %v has no MX records`, e.Domain)
}

// Validate checks that given email address domain is on allow list
func (e ErrDomainNotInAllowedList) Validate(email string) (errs []error) {
	if strings.Count(email, `@`) != 1 {
		return append(errs, &ErrNoAt{})
	}

	parts := strings.Split(email, `@`)

	e.Domain = parts[1]

	for _, d := range e.allowedDomains {
		if e.Domain == d {
			// Found
			return nil
		}
	}

	errs = append(errs, &ErrDomainNotInAllowedList{
		Domain: e.Domain,
		Err:    fmt.Errorf(`not allowed domain: %q`, e.Domain),
	})

	return errs
}
