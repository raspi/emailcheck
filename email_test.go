package emailcheck

import (
	"errors"
	"github.com/raspi/emailcheck/rules"
	"testing"
)

func TestEmpty(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrEmpty{}})
	errs := validator.Validate(``)

	if len(errs) == 0 {
		t.Fatalf(`there should be at least one error`)
	}

	eCount := 0

	for _, e := range errs {
		if errors.Is(e, e.(*rules.ErrEmpty)) {
			eCount++
		}
	}

	if eCount != 1 {
		t.Fail()
	}

}

func TestNoAt(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrNoAt{}})
	errs := validator.Validate(`hello`)

	if len(errs) != 1 {
		t.Fatalf(`there should be at least one error`)
	}

	eCount := 0

	for _, e := range errs {
		if errors.Is(e, e.(*rules.ErrNoAt)) {
			eCount++
		}
	}

	if eCount != 1 {
		t.Fail()
	}

}

func TestLocal(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrEmpty{}, rules.ErrNoAt{}})
	errs := validator.Validate(`user@localhost`)

	if len(errs) != 0 {
		t.Fail()
	}
}

func TestNoDotAfterAt(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrNoDotAfterAt{}})
	errs := validator.Validate(`user@example`)

	if len(errs) != 1 {
		t.Fatalf(`there should be at least one error`)
	}

	eCount := 0

	for _, e := range errs {
		if errors.Is(e, e.(*rules.ErrNoDotAfterAt)) {
			eCount++
		}
	}

	if eCount != 1 {
		t.Fail()
	}

}

func TestDotAfterAt(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrNoDotAfterAt{}})
	errs := validator.Validate(`user@example.org`)

	if len(errs) != 0 {
		t.Fail()
	}
}

func TestNoUser(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrNoUser{}})
	errs := validator.Validate(`@example.org`)

	if len(errs) != 1 {
		t.Fail()
	}
}

func TestHasPrefixSpace(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrHasSpaces{}})
	errs := validator.Validate(` user@example.org`)

	if len(errs) != 1 {
		t.Fail()
	}
}

func TestHasSuffixSpace(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrHasSpaces{}})
	errs := validator.Validate(`user@example.org `)

	if len(errs) != 1 {
		t.Fail()
	}
}

func TestHasSpace(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrHasSpaces{}})
	errs := validator.Validate(`user@ example.org`)

	if len(errs) != 1 {
		t.Fail()
	}
}

func TestHasSpace2(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrHasSpaces{}})
	errs := validator.Validate(`user @example.org`)

	if len(errs) != 1 {
		t.Fail()
	}
}

func TestNoMXDomain(t *testing.T) {
	validator := New([]rules.Validator{rules.ErrNoDomainMXRecords{}})
	errs := validator.Validate(`user@localhost`)

	if len(errs) != 1 {
		t.Fail()
	}

	eCount := 0

	for _, e := range errs {
		if errors.Is(e, e.(*rules.ErrNoDomainMXRecords)) {
			eCount++
		}
	}

	if eCount != 1 {
		t.Fail()
	}

}

func TestDomainNotInAllowedList(t *testing.T) {
	allowedlist := rules.NewErrDomainNotInAllowedList([]string{`localhost`})
	validator := New([]rules.Validator{allowedlist})
	errs := validator.Validate(`user@example.org`)

	if len(errs) != 1 {
		t.Fail()
	}

	eCount := 0

	for _, e := range errs {
		if errors.Is(e, e.(*rules.ErrDomainNotInAllowedList)) {
			eCount++
		}
	}

	if eCount != 1 {
		t.Fail()
	}
}
