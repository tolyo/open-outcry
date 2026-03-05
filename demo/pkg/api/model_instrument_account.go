package api

// InstrumentAccount struct for InstrumentAccount
type InstrumentAccount struct {
	Id          string                     `json:"id"`
	Instruments []InstrumentAccountHolding `json:"instruments"`
}
