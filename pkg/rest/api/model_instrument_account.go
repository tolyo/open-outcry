package api

type InstrumentAccount struct {
	Id string `json:"id"`

	Instruments []InstrumentAccountHolding `json:"instruments"`
}

// AssertInstrumentAccountRequired checks if the required fields are not zero-ed
func AssertInstrumentAccountRequired(obj InstrumentAccount) error {
	elements := map[string]interface{}{
		"id":          obj.Id,
		"instruments": obj.Instruments,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Instruments {
		if err := AssertInstrumentAccountHoldingRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertInstrumentAccountConstraints checks if the values respects the defined constraints
func AssertInstrumentAccountConstraints(obj InstrumentAccount) error {
	return nil
}
