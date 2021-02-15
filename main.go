package emailcheck

import (
	"github.com/raspi/emailcheck/rules"
)

type EmailValidator struct {
	rules []rules.Validator
}

func (ev EmailValidator) Validate(email string) (errs []error) {
	for _, r := range ev.rules {
		errs = append(errs, r.Validate(email)...)
	}

	return errs
}

func New(rules []rules.Validator) EmailValidator {
	return EmailValidator{
		rules: rules,
	}
}
