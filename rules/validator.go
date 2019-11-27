package rules

type Validator interface {
	Validate(e string) []error
}

// Default validation rules
var DefaultRules = []Validator{
	ErrEmpty{},
	ErrNoAt{},
	ErrNoUser{},
	ErrHasSpaces{},
	ErrNoDotAfterAt{},
	ErrNoDomainMXRecords{},
}
