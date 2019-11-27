package rules

import (
	"fmt"
	"net"
	"strings"
)

type ErrNoDomainMXRecords struct {
	Domain string
	Err    error
}

func (e *ErrNoDomainMXRecords) Error() string {
	return fmt.Sprintf(`domain %v has no MX records`, e.Domain)
}

func (e *ErrNoDomainMXRecords) Unwrap() error { return e.Err }

func (e ErrNoDomainMXRecords) Validate(email string) (errs []error) {
	if strings.Count(email, `@`) != 1 {
		return append(errs, &ErrNoAt{})
	}

	parts := strings.Split(email, `@`)

	e.Domain = parts[1]

	mxs, err := net.LookupMX(e.Domain)

	if err != nil {
		errs = append(errs, &ErrNoDomainMXRecords{
			Domain: e.Domain,
			Err:    fmt.Errorf(`no MX: %w`, err),
		})
		return errs
	}

	if len(mxs) == 0 {
		return append(errs, &e)
	}

	return errs
}
