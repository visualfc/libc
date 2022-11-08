package libc

import unsafe "unsafe"

func wcsrtombs(s *int8, ws **int32, n uint64, st *Struct___mbstate_t) uint64 {
	var ws2 *int32
	var buf [4]int8
	var N uint64 = n
	var l uint64
	if !(s != nil) {
		for func() *int32 {
			n = uint64(0)
			return func() (_cgo_ret *int32) {
				_cgo_addr := &ws2
				*_cgo_addr = *ws
				return *_cgo_addr
			}()
		}(); *ws2 != 0; *(*uintptr)(unsafe.Pointer(&ws2)) += 4 {
			if uint32(*ws2) >= uint32(128) {
				l = wcrtomb((*int8)(unsafe.Pointer(&buf)), *ws2, nil)
				if !(l+uint64(1) != 0) {
					return uint64(18446744073709551615)
				}
				n += l
			} else {
				n++
			}
		}
		return n
	}
	for n >= uint64(4) {
		if uint32(**ws)-uint32(1) >= uint32(127) {
			if !(**ws != 0) {
				*s = int8(0)
				*ws = (*int32)(nil)
				return N - n
			}
			l = wcrtomb(s, **ws, nil)
			if !(l+uint64(1) != 0) {
				return uint64(18446744073709551615)
			}
			*(*uintptr)(unsafe.Pointer(&s)) += uintptr(l)
			n -= l
		} else {
			*func() (_cgo_ret *int8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = int8(**ws)
			n--
		}
		*(*uintptr)(unsafe.Pointer(&*ws)) += 4
	}
	for n != 0 {
		if uint32(**ws)-uint32(1) >= uint32(127) {
			if !(**ws != 0) {
				*s = int8(0)
				*ws = (*int32)(nil)
				return N - n
			}
			l = wcrtomb((*int8)(unsafe.Pointer(&buf)), **ws, nil)
			if !(l+uint64(1) != 0) {
				return uint64(18446744073709551615)
			}
			if l > n {
				return N - n
			}
			wcrtomb(s, **ws, nil)
			*(*uintptr)(unsafe.Pointer(&s)) += uintptr(l)
			n -= l
		} else {
			*func() (_cgo_ret *int8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = int8(**ws)
			n--
		}
		*(*uintptr)(unsafe.Pointer(&*ws)) += 4
	}
	return N
}
