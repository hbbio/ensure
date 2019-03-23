package ensure

import (
	"errors"
	"testing"
)

func TestMake(t *testing.T) {
	e := Make(t)
	err := errors.New("error")
	e.Ensure(err).Fails()
	err = nil
	e.Ensure(err).Succeeds()
	s := "foo"
	e.Ensure(s).Is("foo")
	e.Ensure(s).IsNot("bar")
	e.Ensure(s).IsNotEmpty()
	v := 10
	e.Ensure(v).Is(10)
	e.Ensure(v).IsNot(11)
}
