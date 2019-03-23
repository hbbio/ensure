`ensure` is a minimal Go package that eases writing tests.

# Example

```go
func TestXXX(t *testing.T) {
	e := ensure.Make(t)
    ...
	e.Ensure(...some func call...).Fails()
    ...
	e.Ensure(...another func call...).Succeeds()
    ...
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

# Supported types

As of now, `ensure` just support a few types but might be expanded in the future.

# About

Written by @hbbio and released under the MIT license.
