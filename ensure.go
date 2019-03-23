// Package ensure is a minimal Go package that eases writing tests.
package ensure

import (
	"log"
	"testing"
)

// Testable respresents a value to test
type Testable struct {
	Test   *testing.T
	Error  error
	String string
	Value  interface{}
}

// Succeeds expects the Testable (error) to pass
func (t Testable) Succeeds() {
	if t.Error != nil {
		log.Fatalf("%v: fails with error '%v'\n", t.Test.Name(), t.Error)
	}
}

// Fails expects the Testable (error) to fail
func (t Testable) Fails() {
	if t.Error == nil {
		log.Fatalf("%v: should have failed\n", t.Test.Name())
	}
}

// Is expects the Testable to be the same
func (t Testable) Is(v interface{}) {
	switch v.(type) {
	case error:
		log.Fatalf("%v: do not use 'Is' for errors\n", t.Test.Name())
	case string:
		s := v.(string)
		if s != t.String {
			log.Fatalf("%v: should have similar values (is: '%v', expected: '%v')\n", t.Test.Name(), t.String, s)
		}
	default:
		if v != t.Value {
			log.Fatalf("%v: should have similar values (is: '%v', expected: '%v')\n", t.Test.Name(), t.Value, v)
		}
	}
}

// IsNot expect the Testable to be different
func (t Testable) IsNot(v interface{}) {
	switch v.(type) {
	case error:
		log.Fatalf("%v: do not use 'Is' for errors\n", t.Test.Name())
	case string:
		s := v.(string)
		if s == t.String {
			log.Fatalf("%v: should have different values (value: '%v')\n", t.Test.Name(), s)
		}
	default:
		if v == t.Value {
			log.Fatalf("%v: should have different values (value: '%v')\n", t.Test.Name(), v)
		}
	}
}

// IsNotEmpty expect the Testable to be a non-empty string
func (t Testable) IsNotEmpty() {
	if len(t.String) == 0 {
		log.Fatalf("%v: string should not be empty", t.Test.Name())
	}
}

func makeEnsure(v interface{}, t *testing.T) Testable {
	switch v.(type) {
	case error:
		err := v.(error)
		return Testable{Test: t, Error: err}
	case string:
		s := v.(string)
		return Testable{Test: t, String: s}
	default:
		return Testable{Test: t, Value: v}
	}
}

// Ensure creates a Testable result (without testing integration)
func Ensure(v interface{}) Testable {
	return makeEnsure(v, nil)
}

// T represents an Ensure for a test
type T struct {
	Ensure func(v interface{}) Testable
}

// Make returns the Ensure function integrated testing.T
func Make(t *testing.T) T {
	return T{
		Ensure: func(v interface{}) Testable { return makeEnsure(v, t) },
	}
}
