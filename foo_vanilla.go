package main // import "foo"

// #cgo LDFLAGS: ${LDFLAGS}
// #include <stdlib.h>
// #include "c_cpp/foo.h"
import "C"
import "unsafe"

func main() {
	data := "Hello World!!"

	cstr := C.CString(data)
	defer C.free(unsafe.Pointer(cstr))

	C.DoSay(cstr, &cstr)
}
