package libc

import unsafe "unsafe"

func Mbtowc(wc *int32, src *int8, n uint64) int32 {
	var c uint32
	var s *uint8 = (*uint8)(unsafe.Pointer(src))
	var dummy int32
	if !(s != nil) {
		return int32(0)
	}
	if !(n != 0) {
		goto ilseq
	}
	if !(wc != nil) {
		wc = &dummy
	}
	if int32(*s) < int32(128) {
		return func() int32 {
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
		// if !!(*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__pthread_self().Locale.Cat)))) + uintptr(int32(0))*8)) != nil) {
		return int32(4)
		// } else {
		// 	return int32(1)
		// }
	}() == int32(1) {
		return func() int32 {
			*wc = int32(57343) & int32(int8(*s))
			return int32(1)
		}()
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
	if n < uint64(4) && c<<(uint64(6)*n-uint64(6))&2147483648 != 0 {
		goto ilseq
	}
	if (int32(*s)>>int32(3)-int32(16)|(int32(*s)>>int32(3)+int32(c)>>int32(26)))&-8 != 0 {
		goto ilseq
	}
	c = c<<int32(6) | uint32(int32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &s
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}())-int32(128))
	if !(c&2147483648 != 0) {
		*wc = int32(c)
		return int32(2)
	}
	if uint32(*s)-uint32(128) >= uint32(64) {
		goto ilseq
	}
	c = c<<int32(6) | uint32(int32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &s
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}())-int32(128))
	if !(c&2147483648 != 0) {
		*wc = int32(c)
		return int32(3)
	}
	if uint32(*s)-uint32(128) >= uint32(64) {
		goto ilseq
	}
	*wc = int32(c<<int32(6) | uint32(int32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &s
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}())-int32(128)))
	return int32(4)
ilseq:
	*X__errno_location() = int32(84)
	return -1
}
