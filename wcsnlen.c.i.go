package libc

import unsafe "unsafe"

func Wcsnlen(s *int32, n uint64) uint64 {
	var z *int32 = wmemchr(s, int32(0), n)
	if z != nil {
		n = uint64((uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(s))) / 4)
	}
	return n
}
