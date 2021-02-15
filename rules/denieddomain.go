package rules

import (
	"fmt"
	"strings"
)

type ErrDomainInDeniedList struct {
	Domain        string
	Err           error
	deniedDomains []string // list of denied (blocked) domains
}

func NewErrDomainInDeniedList(denied []string) ErrDomainInDeniedList {
	return ErrDomainInDeniedList{
		deniedDomains: denied,
	}
}

func (e *ErrDomainInDeniedList) Error() string {
	return fmt.Sprintf(`domain %q is denied`, e.Domain)
}

// Validate checks that given email address domain is on denied list
func (e ErrDomainInDeniedList) Validate(email string) (errs []error) {
	if strings.Count(email, `@`) != 1 {
		return append(errs, &ErrNoAt{})
	}

	parts := strings.Split(email, `@`)

	e.Domain = parts[1]

	for _, d := range e.deniedDomains {
		if e.Domain == d {
			// Found (denied)
			errs = append(errs, &ErrDomainInDeniedList{
				Domain: e.Domain,
				Err:    fmt.Errorf(`not allowed domain: %q`, e.Domain),
			})

			return errs
		}
	}

	return nil
}
