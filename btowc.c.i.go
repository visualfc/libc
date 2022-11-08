package libc

import unsafe "unsafe"

func btowc(c int32) uint32 {
	var b int32 = int32(uint8(c))
	return func() uint32 {
		if uint32(b) < uint32(128) {
			return uint32(b)
		} else {
			return func() uint32 {
				if func() int32 {
					if !!(*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__pthread_self().Locale.Cat)))) + uintptr(int32(0))*8)) != nil) {
						return int32(4)
					} else {
						return int32(1)
					}
				}() == int32(1) && c != -1 {
					return uint32(int32(57343) & int32(int8(c)))
				} else {
					return uint32(4294967295)
				}
			}()
		}
	}()
}
