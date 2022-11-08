package libc

import unsafe "unsafe"

func wcpcpy(d *int32, s *int32) *int32 {
	return (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(Wcscpy(d, s))) + uintptr(Wcslen(s))*4))
}
