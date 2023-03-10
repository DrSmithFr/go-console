package validator

type Validator func(string) error

func MakeChainedValidator(validators ...Validator) Validator {
	return func(answer string) error {
		for _, validator := range validators {
			if err := validator(answer); err != nil {
				return err
			}
		}

		return nil
	}
}
