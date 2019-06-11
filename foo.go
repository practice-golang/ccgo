package main // import "gofoo"

// #cgo LDFLAGS: -L./ -lfoo -lstdc++
// #include "c_cpp/foo.h"
import "C"

func main() {
	C.DoSay()
}
