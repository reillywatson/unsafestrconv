package a

import "fmt"

func Foo(x int) string {
	return string(x) // want "Unsafe string\\(int\\) conversion"
}

const c int = 5

func Bar() {
	fmt.Println(string(c)) // want "Unsafe string\\(int\\) conversion"
}

func Baz() {
	var z rune = 'a'
	fmt.Println(string(z))
}

func Qux() {
	var a byte = byte(c)
	fmt.Println(string(a))
}

type MyInt int

func WithMyInt() {
	var x MyInt = 5
	fmt.Println(string(x)) // want "Unsafe string\\(int\\) conversion"
}
