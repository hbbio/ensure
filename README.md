`ensure` is a minimal Go package that eases writing tests.

[![GoDoc](https://godoc.org/github.com/hbbio/ensure?status.svg)](https://godoc.org/github.com/hbbio/ensure)
[![Build
Status](https://travis-ci.org/hbbio/ensure.svg?branch=master)](https://travis-ci.org/hbbio/ensure)

# Example

```go
func TestXXX(t *testing.T) {
	e := ensure.Make(t)
    // ...
	e.Ensure(...some func call...).Fails()
    // ...
	e.Ensure(...another func call...).Succeeds()
    // ...
    e.Ensure(...).Is(value)
}
```

# Why

I wanted a minimal (no dependencies) way to avoid writing long test files that need to create their own error messages and that maximize readability in a single line.

Write:

```go
e.Ensure(...some func call...).Fails()
```

instead of

```go
err := ...some func call...
if err == nil {
    log.Fatalf("this should have failed (test %v)", t.Name())
}
```

[gomega](https://onsi.github.io/gomega/) is a more full-featured assertion library but is too big for me and has too many dependencies (including the whole Ginkgo framework).

# Usage without `testing` integration

You can also directly call `ensure.Ensure` without calling `Make`.

# Supported types and checks

As of now, `ensure` just support a few types but might be expanded in the future.

| Type                | Checks                  |
| ------------------- | ----------------------- |
| error               | Succeeds(), Fails()     |
| string              | IsNotEmpty()            |
| anything but errors | Is(value), IsNot(value) |

# About

Written by [@hbbio](https://github.com/hbbio) and released under the MIT license.
