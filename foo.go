package main // import "gofoo"

// If build without libfoo.a, disable below LDFLAGS comment

// #cgo LDFLAGS: -L./ -lfoo -lstdc++
// #include "foo.h"
import "C"

type goFoo struct {
	foo C.Foo
}

func new() goFoo {
	var ret goFoo
	ret.foo = C.FooInit()
	return ret
}

func (f goFoo) Free() {
	C.FooFree(f.foo)
}

func (f goFoo) Bar() {
	C.FooBar(f.foo)
}

func main() {
	foo := new()
	foo.Bar()
	foo.Free()
}
