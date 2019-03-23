`ensure` is a minimal Go package that eases writing tests.

[![GoDoc](https://godoc.org/github.com/hbbio/ensure?status.svg)](https://godoc.org/github.com/hbbio/ensure)
[![Build
Status](https://travis-ci.org/hbbio/ensure.svg?branch=master)](https://travis-ci.org/hbbio/ensure)

## Examples

```go
func TestXXX(t *testing.T) {
  e := ensure.Make(t)
  e.Ensure(os.Remove("/etc/pAsswd")).Fails()
  e.Ensure(os.Setenv("ENSURE", "ISCOOL")).Succeeds()
  s := "foo"
  e.Ensure(s).Is("foo")
  e.Ensure(s).IsNot("bar")
}
```

## Why

I wanted a minimal (no dependencies) way to avoid writing long test files that need to create their own error messages and that maximize readability in a single line.

Write:

```go
e.Ensure(someFunc(xxx)).Fails()
```

instead of

```go
err := someFunc(xxx)
if err == nil {
    log.Fatalf("this should have failed (test %v)", t.Name())
}
```

You can also directly call `ensure.Ensure` without calling `Make`.

## Alternatives

There are several packages, which are full-featured assertion libraries, that do more things but with (much) more code and dependencies:

- [gomega](https://onsi.github.io/gomega/)
- [testify](https://github.com/stretchr/testify)
- [check](https://github.com/go-check/check)

## Installation

```sh
go get github.com/hbbio/ensure
```

## Supported types and checks

As of now, `ensure` only supports a few types but might be expanded in the future.

| Type                | Checks                  |
| ------------------- | ----------------------- |
| error               | Succeeds(), Fails()     |
| string              | IsNotEmpty()            |
| anything but errors | Is(value), IsNot(value) |

## About

Written by [@hbbio](https://github.com/hbbio) and released under the MIT license.
