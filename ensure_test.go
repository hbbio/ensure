package ensure

import (
	"errors"
	"os"
	"strconv"
	"testing"
)

func TestDirect(t *testing.T) {
	Ensure(t, nil).Succeeds("ensure succeeds")
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		log.Printf("recovered: %v", r)
	// 	}
	// }()
	// Ensure(t, nil).Fails("ensure fails")
}

func TestMake(t *testing.T) {
	e := Make(t)
	e.Ensure(nil).Succeeds()
}

func TestErrors(t *testing.T) {
	e := Make(t)
	err := errors.New("error")
	e.Ensure(err).Fails()
	err = nil
	e.Ensure(err).Succeeds()
}

func TestString(t *testing.T) {
	e := Make(t)
	s := "foo"
	e.Ensure(s).Is("foo")
	e.Ensure("foo").Is("foo")
	e.Ensure(s).IsNot("bar")
	e.Ensure(s).Contains("fo")
	s2 := s
	e.Ensure(s).Is(s2)
	e.Ensure(s).IsNotEmpty()
}

func TestInt(t *testing.T) {
	e := Make(t)
	v := 10
	e.Ensure(v).Is(10).IsNot(11)
}

func TestEnsure2(t *testing.T) {
	e := Make(t)
	e.Ensure2(os.Open("/etc/passwd")).Succeeds("in opening /etc/passwd")
	b := e.Ensure2(strconv.ParseBool("true")).Succeeds().Return()
	e.Ensure(b).Is(true)
}
