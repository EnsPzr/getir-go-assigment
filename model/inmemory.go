package model

type InMemory struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var (
	keyLenghtError   = "Key can be minimum 2 character and maximum 100 character"
	valueLenghtError = "Value can be minimum 2 character and maximum 100 character"
)

// Validate
// This method validates the model where comes from in-memory and in-memory-cache.
func (i *InMemory) Validate() []string {
	errs := make([]string, 0)
	if len(i.Key) < 2 || len(i.Key) > 100 {
		errs = append(errs, keyLenghtError)
	}
	if len(i.Value) < 2 || len(i.Value) > 100 {
		errs = append(errs, valueLenghtError)
	}
	return errs
}
