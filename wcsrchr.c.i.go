package libc

import unsafe "unsafe"

func Wcsrchr(s *int32, c int32) *int32 {
	var p *int32
	for p = (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(Wcslen(s))*4)); uintptr(unsafe.Pointer(p)) >= uintptr(unsafe.Pointer(s)) && *p != c; *(*uintptr)(unsafe.Pointer(&p)) -= 4 {
	}
	return func() *int32 {
		if uintptr(unsafe.Pointer(p)) >= uintptr(unsafe.Pointer(s)) {
			return (*int32)(unsafe.Pointer(p))
		} else {
			return nil
		}
	}()
}
