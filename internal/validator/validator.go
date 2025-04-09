package validator

import "fmt"

// Validator is a struct for holding validation errors.
type Validator struct {
	Errors []string
}

// New creates a new Validator instance.
func New() *Validator {
	return &Validator{}
}

// Check validates a condition and stores an error message if the condition fails.
func (v *Validator) Check(condition bool, field, message string) {
	if !condition {
		v.Errors = append(v.Errors, fmt.Sprintf("%s: %s", field, message))
	}
}

// Valid checks if there are no validation errors.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}
