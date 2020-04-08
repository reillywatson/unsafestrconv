# unsafestrconv
An analyzer to detect string(int) casts in Go programs

-[![CircleCI](https://circleci.com/gh/reillywatson/unsafestrconv.svg?style=svg)](https://circleci.com/gh/reillywatson/unsafestrconv)
-[![codecov](https://codecov.io/gh/reillywatson/unsafestrconv/branch/master/graph/badge.svg)](https://codecov.io/gh/reillywatson/unsafestrconv)

Constructs like this are legal in Go, but probably programming errors:

```go
x := 100
fmt.Println(string(x)) // outputs "d", not "100"
```

This linter helps catch those!

Note that the Go stdlib now has a better version of this check, you should use it instead! See https://github.com/golang/tools/tree/master/go/analysis/passes/stringintconv/cmd/stringintconv.
