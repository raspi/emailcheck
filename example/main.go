package main

import (
	"fmt"
	"github.com/raspi/emailcheck"
	"github.com/raspi/emailcheck/rules"
)

func main() {
	validator := emailcheck.New(rules.DefaultRules)
	errs := validator.Validate("email@address.tld")

	if len(errs) == 0 {
		// No errors
		return
	}

	for _, err := range errs {
		fmt.Printf("invalid email: %v", err)
	}
}
