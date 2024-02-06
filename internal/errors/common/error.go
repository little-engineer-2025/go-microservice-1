package presenter

import "fmt"

// NewNil create an error for a nil value at the presenter
// component level.
func NewNil(fieldName string) error {
	if fieldName == "" {
		panic("'fieldName' is empty")
	}
	return fmt.Errorf("'%s' is nil", fieldName)
}

// NewEmpty create an error for an empty value at the
// presenter level.
func NewEmpty(fieldName string) error {
	if fieldName == "" {
		panic("'fieldName' is empty")
	}
	return fmt.Errorf("'%s' is empty", fieldName)
}
