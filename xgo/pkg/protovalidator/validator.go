package protovalidator

// Validator is a general interface that allows a message to be validated.
type Validator interface {
	Validate() error
}

// InvokeValidatorIfExists for invoke the Validate method if a interface is a Validator.
func InvokeValidatorIfExists(candidate interface{}) error {
	if candidate == nil {
		return nil
	}
	if validator, ok := candidate.(Validator); ok {
		return validator.Validate()
	}
	return nil
}
