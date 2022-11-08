package libc

import unsafe "unsafe"

func Wcsspn(s *int32, c *int32) uint64 {
	var a *int32
	for a = s; *s != 0 && Wcschr(c, *s) != nil; *(*uintptr)(unsafe.Pointer(&s)) += 4 {
	}
	return uint64((uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer(a))) / 4)
}
