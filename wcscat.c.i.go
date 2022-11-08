package libc

import unsafe "unsafe"

func Wcscat(dest *int32, src *int32) *int32 {
	Wcscpy((*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(dest))+uintptr(Wcslen(dest))*4)), src)
	return dest
}
