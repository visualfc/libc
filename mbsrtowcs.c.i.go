package libc

import unsafe "unsafe"

func mbsrtowcs(ws *int32, src **int8, wn uint64, st *Struct___mbstate_t) uint64 {

	var s *uint8 = (*uint8)(unsafe.Pointer(*src))
	var wn0 uint64 = wn
	var c uint32 = uint32(0)
	if !(st != nil && func() (_cgo_ret uint32) {
		_cgo_addr := &c
		*_cgo_addr = *(*uint32)(unsafe.Pointer(st))
		return *_cgo_addr
	}() != 0) {
		goto _cgol_1
	}
	if !(ws != nil) {
		goto _cgol_3
	}
	*(*uint32)(unsafe.Pointer(st)) = uint32(0)
	goto resume
	goto _cgol_2
_cgol_3:
	goto resume0
_cgol_2:
	;
_cgol_1:
	if func() int32 {
		if !!(*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__pthread_self().Locale.Cat)))) + uintptr(int32(0))*8)) != nil) {
			return int32(4)
		} else {
			return int32(1)
		}
	}() == int32(1) {
		if !(ws != nil) {
			return Strlen((*int8)(unsafe.Pointer(s)))
		}
		for {
			if !(wn != 0) {
				*src = (*int8)(unsafe.Pointer(s))
				return wn0
			}
			if !(*s != 0) {
				break
			}
			c = uint32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
			*func() (_cgo_ret *int32) {
				_cgo_addr := &ws
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = int32(57343) & int32(int8(c))
			wn--
		}
		*ws = int32(0)
		*src = (*int8)(nil)
		return wn0 - wn
	}
	if !!(ws != nil) {
		goto _cgol_5
	}
_cgol_6:
	if uint32(*s)-uint32(1) < uint32(127) && uint64(uintptr(unsafe.Pointer(s)))%uint64(4) == uint64(0) {
		for !((*(*uint32)(unsafe.Pointer(s))|(*(*uint32)(unsafe.Pointer(s))-uint32(16843009)))&uint32(2155905152) != 0) {
			*(*uintptr)(unsafe.Pointer(&s)) += uintptr(int32(4))
			wn -= uint64(4)
		}
	}
	if uint32(*s)-uint32(1) < uint32(127) {
		*(*uintptr)(unsafe.Pointer(&s))++
		wn--
		goto _cgol_6
	}
	if uint32(*s)-uint32(194) > 50 {
		goto _cgol_7
	}
	c = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&__fsmu8)))) + uintptr(uint32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &s
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}())-uint32(194))*4))
resume0:
	if (int32(*s)>>int32(3)-int32(16)|(int32(*s)>>int32(3)+int32(c)>>int32(26)))&-8 != 0 {
		*(*uintptr)(unsafe.Pointer(&s))--
		goto _cgol_7
	}
	*(*uintptr)(unsafe.Pointer(&s))++
	if c&33554432 != 0 {
		if uint32(*s)-uint32(128) >= uint32(64) {
			*(*uintptr)(unsafe.Pointer(&s)) -= uintptr(int32(2))
			goto _cgol_7
		}
		*(*uintptr)(unsafe.Pointer(&s))++
		if c&524288 != 0 {
			if uint32(*s)-uint32(128) >= uint32(64) {
				*(*uintptr)(unsafe.Pointer(&s)) -= uintptr(int32(3))
				goto _cgol_7
			}
			*(*uintptr)(unsafe.Pointer(&s))++
		}
	}
	wn--
	c = uint32(0)
	goto _cgol_6
_cgol_7:
	goto _cgol_4
_cgol_5:
	;
_cgol_8:
	if !(wn != 0) {
		*src = (*int8)(unsafe.Pointer(s))
		return wn0
	}
	if uint32(*s)-uint32(1) < uint32(127) && uint64(uintptr(unsafe.Pointer(s)))%uint64(4) == uint64(0) {
		for wn >= uint64(5) && !((*(*uint32)(unsafe.Pointer(s))|(*(*uint32)(unsafe.Pointer(s))-uint32(16843009)))&uint32(2155905152) != 0) {
			*func() (_cgo_ret *int32) {
				_cgo_addr := &ws
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
			*func() (_cgo_ret *int32) {
				_cgo_addr := &ws
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
			*func() (_cgo_ret *int32) {
				_cgo_addr := &ws
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
			*func() (_cgo_ret *int32) {
				_cgo_addr := &ws
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
			wn -= uint64(4)
		}
	}
	if uint32(*s)-uint32(1) < uint32(127) {
		*func() (_cgo_ret *int32) {
			_cgo_addr := &ws
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() = int32(*func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}())
		wn--
		goto _cgol_8
	}
	if uint32(*s)-uint32(194) > 50 {
		goto _cgol_9
	}
	c = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&__fsmu8)))) + uintptr(uint32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &s
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}())-uint32(194))*4))
resume:
	if (int32(*s)>>int32(3)-int32(16)|(int32(*s)>>int32(3)+int32(c)>>int32(26)))&-8 != 0 {
		*(*uintptr)(unsafe.Pointer(&s))--
		goto _cgol_9
	}
	c = c<<int32(6) | uint32(int32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &s
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}())-int32(128))
	if c&2147483648 != 0 {
		if uint32(*s)-uint32(128) >= uint32(64) {
			*(*uintptr)(unsafe.Pointer(&s)) -= uintptr(int32(2))
			goto _cgol_9
		}
		c = c<<int32(6) | uint32(int32(*func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}())-int32(128))
		if c&2147483648 != 0 {
			if uint32(*s)-uint32(128) >= uint32(64) {
				*(*uintptr)(unsafe.Pointer(&s)) -= uintptr(int32(3))
				goto _cgol_9
			}
			c = c<<int32(6) | uint32(int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())-int32(128))
		}
	}
	*func() (_cgo_ret *int32) {
		_cgo_addr := &ws
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
		return
	}() = int32(c)
	wn--
	c = uint32(0)
	goto _cgol_8
_cgol_9:
	;
_cgol_4:
	if !(c != 0) && !(*s != 0) {
		if ws != nil {
			*ws = int32(0)
			*src = (*int8)(nil)
		}
		return wn0 - wn
	}
	*X__errno_location() = int32(84)
	if ws != nil {
		*src = (*int8)(unsafe.Pointer(s))
	}
	return uint64(18446744073709551615)
}
