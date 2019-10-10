# unsafestrconv
An analyzer to detect string(int) casts in Go programs

Constructs like this are legal in Go, but probably programming errors:
x := 100
fmt.Println(string(x)) // outputs "d", not "100"

This linter helps catch those!