package libc

import unsafe "unsafe"

func wcpncpy(d *int32, s *int32, n uint64) *int32 {
	return (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(Wcsncpy(d, s, n))) + uintptr(Wcsnlen(s, n))*4))
}
