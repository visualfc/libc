package libc

import unsafe "unsafe"

func mbrtowc(wc *int32, src *int8, n uint64, st *Struct___mbstate_t) uint64 {
	var c uint32
	var s *uint8 = (*uint8)(unsafe.Pointer(src))
	var N uint32 = uint32(n)
	var dummy int32
	if !(st != nil) {
		st = (*Struct___mbstate_t)(unsafe.Pointer(&_cgos_mbrtowc_internal_state_mbrtowc))
	}
	c = *(*uint32)(unsafe.Pointer(st))
	if !(s != nil) {
		if c != 0 {
			goto ilseq
		}
		return uint64(0)
	} else if !(wc != nil) {
		wc = &dummy
	}
	if !(n != 0) {
		return uint64(18446744073709551614)
	}
	if !(c != 0) {
		if int32(*s) < int32(128) {
			return func() uint64 {
				if !!(func() (_cgo_ret int32) {
					_cgo_addr := &*wc
					*_cgo_addr = int32(*s)
					return *_cgo_addr
				}() != 0) {
					return 1
				} else {
					return 0
				}
			}()
		}
		if func() int32 {
			if !!(*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__pthread_self().Locale.Cat)))) + uintptr(int32(0))*8)) != nil) {
				return int32(4)
			} else {
				return int32(1)
			}
		}() == int32(1) {
			return uint64(func() int32 {
				*wc = int32(57343) & int32(int8(*s))
				return int32(1)
			}())
		}
		if uint32(*s)-uint32(194) > 50 {
			goto ilseq
		}
		c = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&__fsmu8)))) + uintptr(uint32(*func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}())-uint32(194))*4))
		n--
	}
	if n != 0 {
		if (int32(*s)>>int32(3)-int32(16)|(int32(*s)>>int32(3)+int32(c)>>int32(26)))&-8 != 0 {
			goto ilseq
		}
	loop:
		c = c<<int32(6) | uint32(int32(*func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}())-int32(128))
		n--
		if !(c&2147483648 != 0) {
			*(*uint32)(unsafe.Pointer(st)) = uint32(0)
			*wc = int32(c)
			return uint64(N) - n
		}
		if n != 0 {
			if uint32(*s)-uint32(128) >= uint32(64) {
				goto ilseq
			}
			goto loop
		}
	}
	*(*uint32)(unsafe.Pointer(st)) = c
	return uint64(18446744073709551614)
ilseq:
	*(*uint32)(unsafe.Pointer(st)) = uint32(0)
	*X__errno_location() = int32(84)
	return uint64(18446744073709551615)
}

var _cgos_mbrtowc_internal_state_mbrtowc uint32
