package api

// NewInterface creates a new empty interface with a given value.
func NewInterface(value interface{}) *interface{} {
	v := value
	return &v
}
