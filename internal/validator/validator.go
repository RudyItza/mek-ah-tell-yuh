package validator

type Validator struct {
	Errors map[string][]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string][]string),
	}
}

func (v *Validator) Check(value bool, field, message string) {
	if !value {
		v.Errors[field] = append(v.Errors[field], message)
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}
