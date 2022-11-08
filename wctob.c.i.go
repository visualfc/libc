package libc

import unsafe "unsafe"

func wctob(c uint32) int32 {
	if c < uint32(128) {
		return int32(c)
	}
	if func() int32 {
		if !!(*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__pthread_self().Locale.Cat)))) + uintptr(int32(0))*8)) != nil) {
			return int32(4)
		} else {
			return int32(1)
		}
	}() == int32(1) && uint32(c)-uint32(57216) < uint32(128) {
		return int32(uint8(c))
	}
	return -1
}
