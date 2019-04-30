// Package ensure is a minimal Go package that eases writing tests.
package ensure

import (
	"fmt"
	"log"
	"testing"
)

// Testable respresents a value to test.
type Testable struct {
	Test   *testing.T
	Error  error
	String string
	Value  interface{}
}

// func s(f string, v ...interface{}) string {
// 	return fmt.Sprintf(f, v)
// }
var s = fmt.Sprintf

// Fatal stops with fatal error.
func (t *Testable) Fatal(msg string, what []string) {
	if len(what) > 0 {
		log.Println(what)
	}
	t.Test.Fatal(s("%s: %s\n", t.Test.Name(), msg))
}

// Succeeds expects the Testable (error) to pass.
func (t *Testable) Succeeds(what ...string) *Testable {
	if t.Error != nil {
		t.Fatal(s("fails with error '%v'", t.Error), what)
	}
	return t
}

// Fails expects the Testable (error) to fail.
func (t *Testable) Fails(what ...string) *Testable {
	if t.Error == nil {
		t.Fatal("should have failed", what)
	}
	return t
}

// Is expects the Testable to be the same.
func (t *Testable) Is(v interface{}, what ...string) *Testable {
	switch v.(type) {
	case error:
		t.Fatal("do not use 'Is' for errors", what)
	case string:
		str := v.(string)
		if str != t.String {
			t.Fatal(s("should have similar values (is: '%v', expected: '%v')", t.String, str), what)
		}
	default:
		if v != t.Value {
			t.Fatal(s("should have similar values (is: '%v', expected: '%v')", t.Value, v), what)
		}
	}
	return t
}

// IsNot expect the Testable to be different.
func (t *Testable) IsNot(v interface{}, what ...string) *Testable {
	switch v.(type) {
	case error:
		t.Fatal("do not use 'Is' for errors", what)
	case string:
		str := v.(string)
		if str == t.String {
			t.Fatal(s("should have different values (value: '%v')", str), what)
		}
	default:
		if v == t.Value {
			t.Fatal(s("should have different values (value: '%v')", v), what)
		}
	}
	return t
}

// IsNotEmpty expect the Testable to be a non-empty string.
func (t *Testable) IsNotEmpty(what ...string) *Testable {
	if len(t.String) == 0 {
		t.Fatal("string should not be empty", what)
	}
	return t
}

// makeEnsure constructs an ensure Testable.
func makeEnsure(v interface{}, t *testing.T) *Testable {
	switch v.(type) {
	case error:
		err := v.(error)
		return &Testable{Test: t, Error: err}
	case string:
		s := v.(string)
		return &Testable{Test: t, String: s}
	default:
		return &Testable{Test: t, Value: v}
	}
}

// Ensure creates a Testable result (without testing integration).
func Ensure(v interface{}) *Testable {
	return makeEnsure(v, nil)
}

// T represents an Ensure for a test.
type T struct {
	Ensure  func(v interface{}) *Testable
	Ensure2 func(res, v interface{}) (interface{}, *Testable)
}

// Make returns the Ensure function integrated with testing.
func Make(t *testing.T) T {
	return T{
		Ensure:  func(v interface{}) *Testable { return makeEnsure(v, t) },
		Ensure2: func(res, v interface{}) (interface{}, *Testable) { return res, makeEnsure(v, t) },
	}
}
