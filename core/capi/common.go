package capi

import (
    "unsafe"
    "C"
)

func unsafeStringData(s string) (*C.char, C.int) {
    return *(**C.char)(unsafe.Pointer(&s)), C.int(len(s))
}