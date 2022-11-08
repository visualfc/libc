package libc

import unsafe "unsafe"

func wcrtomb(s *int8, wc int32, st *Struct___mbstate_t) uint64 {
	if !(s != nil) {
		return uint64(1)
	}
	if uint32(wc) < uint32(128) {
		*s = int8(wc)
		return uint64(1)
	} else if func() int32 {
		if !!(*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__pthread_self().Locale.Cat)))) + uintptr(int32(0))*8)) != nil) {
			return int32(4)
		} else {
			return int32(1)
		}
	}() == int32(1) {
		if !(uint32(wc)-uint32(57216) < uint32(128)) {
			*X__errno_location() = int32(84)
			return uint64(18446744073709551615)
		}
		*s = int8(wc)
		return uint64(1)
	} else if uint32(wc) < uint32(2048) {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(int32(192) | wc>>int32(6))
		*s = int8(int32(128) | wc&int32(63))
		return uint64(2)
	} else if uint32(wc) < uint32(55296) || uint32(wc)-uint32(57344) < uint32(8192) {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(int32(224) | wc>>int32(12))
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(int32(128) | wc>>int32(6)&int32(63))
		*s = int8(int32(128) | wc&int32(63))
		return uint64(3)
	} else if uint32(wc)-uint32(65536) < uint32(1048576) {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(int32(240) | wc>>int32(18))
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(int32(128) | wc>>int32(12)&int32(63))
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(int32(128) | wc>>int32(6)&int32(63))
		*s = int8(int32(128) | wc&int32(63))
		return uint64(4)
	}
	*X__errno_location() = int32(84)
	return uint64(18446744073709551615)
}
