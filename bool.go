// Code generated by go generate
// This file was generated by robots at 2018-03-25 18:09:18.750597101 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Bool is an optional bool
type Bool struct {
	value *bool
}

// NewBool creates a optional.Bool from a bool
func NewBool(v bool) Bool {
	return Bool{&v}
}

// Set sets the bool value
func (b Bool) Set(v bool) {
	b.value = &v
}

// Get returns the bool value or an error if not present
func (b Bool) Get() (bool, error) {
	if !b.Present() {
		return *b.value, errors.New("value not present")
	}
	return *b.value, nil
}

// Present returns whether or not the value is present
func (b Bool) Present() bool {
	return b.value != nil
}

// OrElse returns the bool value or a default value if the value is not present
func (b Bool) OrElse(v bool) bool {
	if b.Present() {
		return *b.value
	}
	return v
}

// If calls the function f with the value if the value is present
func (b Bool) If(fn func(bool)) {
	if b.Present() {
		fn(*b.value)
	}
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b.Present() {
		return json.Marshal(b.value)
	}
	return nil, nil
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	if len(data) < 1 {
		b.value = nil
		return nil
	}

	b.value = new(bool)
	return json.Unmarshal(data, b.value)
}
