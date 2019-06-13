package main // import "foo"

// #cgo LDFLAGS: ${LDFLAGS}
// #include <stdlib.h>
// #include "c_cpp/foo.h"
import "C"
import "unsafe"

func main() {
	data1 := "Hello World!!"
	data2 := []string{"Hello", "World!!"}

	cstr1 := C.CString(data1)
	var cstr2 [2]*C.char
	cstr3 := []*C.char{}

	for i, s := range data2 {
		cs := C.CString(s)
		defer C.free(unsafe.Pointer(cs))
		cstr2[i] = cs
		cstr3 = append(cstr3, cs)
	}

	args := (**C.char)(unsafe.Pointer(&cstr2))
	// Slice not work as C pointer so use array instead of slice
	// args := (**C.char)(unsafe.Pointer(&cstr3))

	C.DoSay(cstr1, args)
}
