package libc

import unsafe "unsafe"

func Wcschr(s *int32, c int32) *int32 {
	if !(c != 0) {
		return (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(s)))) + uintptr(Wcslen(s))*4))
	}
	for ; *s != 0 && *s != c; *(*uintptr)(unsafe.Pointer(&s)) += 4 {
	}
	return func() *int32 {
		if *s != 0 {
			return (*int32)(unsafe.Pointer(s))
		} else {
			return nil
		}
	}()
}
