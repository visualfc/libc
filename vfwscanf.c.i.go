package libc

import unsafe "unsafe"

func _cgos_store_int_vfwscanf(dest unsafe.Pointer, size int32, i uint64) {
	if !(dest != nil) {
		return
	}
	switch size {
	case -2:
		*(*int8)(dest) = int8(i)
		break
	case -1:
		*(*int16)(dest) = int16(i)
		break
	case int32(0):
		*(*int32)(dest) = int32(i)
		break
	case int32(1):
		*(*int64)(dest) = int64(i)
		break
	case int32(3):
		*(*int64)(dest) = int64(i)
		break
	}
}
func _cgos_arg_n_vfwscanf(ap []interface {
}, n uint32) unsafe.Pointer {
	var p unsafe.Pointer
	var i uint32
	var ap2 []interface {
	}
	ap2 = ap
	for i = n; i > uint32(1); i-- {
		func(__cgo_args []interface {
		}) (_cgo_ret unsafe.Pointer) {
			_cgo_ret = unsafe.Pointer((*[2]unsafe.Pointer)(unsafe.Pointer(&__cgo_args[0]))[1])
			ap2 = __cgo_args[1:]
			return
		}(ap2)
	}
	p = func(__cgo_args []interface {
	}) (_cgo_ret unsafe.Pointer) {
		_cgo_ret = unsafe.Pointer((*[2]unsafe.Pointer)(unsafe.Pointer(&__cgo_args[0]))[1])
		ap2 = __cgo_args[1:]
		return
	}(ap2)
	return p
}
func _cgos_in_set_vfwscanf(set *int32, c int32) int32 {
	var j int32
	var p *int32 = set
	if *p == '-' {
		if c == '-' {
			return int32(1)
		}
		*(*uintptr)(unsafe.Pointer(&p)) += 4
	} else if *p == ']' {
		if c == ']' {
			return int32(1)
		}
		*(*uintptr)(unsafe.Pointer(&p)) += 4
	}
	for ; *p != 0 && *p != ']'; *(*uintptr)(unsafe.Pointer(&p)) += 4 {
		if *p == '-' && *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4)) != 0 && *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4)) != ']' {
			for j = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(func() (_cgo_ret *int32) {
				_cgo_addr := &p
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}())) - uintptr(1)*4)); j < *p; j++ {
				if c == j {
					return int32(1)
				}
			}
		}
		if c == *p {
			return int32(1)
		}
	}
	return int32(0)
}
func vfwscanf(f *Struct__IO_FILE, fmt *int32, ap []interface {
}) int32 {
	var (
		_tag_cgo1                           int32
		_nm_cgo2                            bool
		_tag_cgo3                           int32
		_nm_cgo4                            bool
		_cgos_vfwscanf_spaces_vfwscanf_cgo5 [22]int32
		gotmatch_cgo6                       int32
	)
	var width int32
	var size int32
	var alloc int32
	var p *int32
	var c int32
	var t int32
	var s *int8
	var wcs *int32
	var dest unsafe.Pointer = unsafe.Pointer(nil)
	var invert int32
	var matches int32 = int32(0)
	var pos int64 = int64(0)
	var cnt int64
	var tmp [22]int8
	var set *int32
	var i uint64
	var k uint64
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	fwide(f, int32(1))
	p = fmt
_cgol_1:
	if !(*p != 0) {
		goto _cgol_2
	}
	alloc = int32(0)
	if iswspace(uint32(*p)) != 0 {
		for iswspace(uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4)))) != 0 {
			*(*uintptr)(unsafe.Pointer(&p)) += 4
		}
		for iswspace(uint32(func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = int32(func() uint32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) && int32(*f.Rpos) < int32(128) {
					return uint32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return getwc(f)
				}
			}())
			return *_cgo_addr
		}())) != 0 {
			pos++
		}
		func() uint32 {
			if f.Rend != nil && uint32(c) < uint32(128) {
				return uint32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return *_cgo_addr
				}())
			} else {
				return ungetwc(uint32(c), f)
			}
		}()
		goto _cgol_1
	}
	if !(*p != '%' || *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4)) == '%') {
		goto _cgol_3
	}
	if *p == '%' {
		*(*uintptr)(unsafe.Pointer(&p)) += 4
		for iswspace(uint32(func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = int32(func() uint32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) && int32(*f.Rpos) < int32(128) {
					return uint32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return getwc(f)
				}
			}())
			return *_cgo_addr
		}())) != 0 {
			pos++
		}
	} else {
		c = int32(func() uint32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) && int32(*f.Rpos) < int32(128) {
				return uint32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return getwc(f)
			}
		}())
	}
	if !(c != *p) {
		goto _cgol_4
	}
	func() uint32 {
		if f.Rend != nil && uint32(c) < uint32(128) {
			return uint32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return *_cgo_addr
			}())
		} else {
			return ungetwc(uint32(c), f)
		}
	}()
	if c < int32(0) {
		goto input_fail
	}
	goto match_fail
_cgol_4:
	pos++
	goto _cgol_1
_cgol_3:
	*(*uintptr)(unsafe.Pointer(&p)) += 4
	if *p == '*' {
		dest = unsafe.Pointer(nil)
		*(*uintptr)(unsafe.Pointer(&p)) += 4
	} else if func() int32 {
		if int32(0) != 0 {
			return iswdigit(uint32(*p))
		} else {
			return func() int32 {
				if uint32(*p)-uint32('0') < uint32(10) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0 && *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))*4)) == '$' {
		dest = _cgos_arg_n_vfwscanf(ap, uint32(*p-'0'))
		*(*uintptr)(unsafe.Pointer(&p)) += uintptr(int32(2)) * 4
	} else {
		dest = func(__cgo_args []interface {
		}) (_cgo_ret unsafe.Pointer) {
			_cgo_ret = unsafe.Pointer((*[2]unsafe.Pointer)(unsafe.Pointer(&__cgo_args[0]))[1])
			ap = __cgo_args[1:]
			return
		}(ap)
	}
	for width = int32(0); func() int32 {
		if int32(0) != 0 {
			return iswdigit(uint32(*p))
		} else {
			return func() int32 {
				if uint32(*p)-uint32('0') < uint32(10) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0; *(*uintptr)(unsafe.Pointer(&p)) += 4 {
		width = int32(10)*width + *p - '0'
	}
	if *p == 'm' {
		wcs = (*int32)(nil)
		s = (*int8)(nil)
		alloc = func() int32 {
			if !!(dest != nil) {
				return 1
			} else {
				return 0
			}
		}()
		*(*uintptr)(unsafe.Pointer(&p)) += 4
	} else {
		alloc = int32(0)
	}
	size = int32(0)
	_tag_cgo1, _nm_cgo2 = *func() (_cgo_ret *int32) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
		return
	}(), true
	if _nm_cgo2 && _tag_cgo1 != 'h' {
		goto _cgol_5
	}
	_nm_cgo2 = false
	if *p == 'h' {
		func() int32 {
			*(*uintptr)(unsafe.Pointer(&p)) += 4
			return func() (_cgo_ret int32) {
				_cgo_addr := &size
				*_cgo_addr = -2
				return *_cgo_addr
			}()
		}()
	} else {
		size = -1
	}
	goto _cgol_6
_cgol_5:
	if _nm_cgo2 && _tag_cgo1 != 'l' {
		goto _cgol_7
	}
	_nm_cgo2 = false
	if *p == 'l' {
		func() int32 {
			*(*uintptr)(unsafe.Pointer(&p)) += 4
			return func() (_cgo_ret int32) {
				_cgo_addr := &size
				*_cgo_addr = int32(3)
				return *_cgo_addr
			}()
		}()
	} else {
		size = int32(1)
	}
	goto _cgol_6
_cgol_7:
	if _nm_cgo2 && _tag_cgo1 != 'j' {
		goto _cgol_8
	}
	_nm_cgo2 = false
	size = int32(3)
	goto _cgol_6
_cgol_8:
	if _nm_cgo2 && _tag_cgo1 != 'z' {
		goto _cgol_9
	}
	_nm_cgo2 = false
_cgol_9:
	if _nm_cgo2 && _tag_cgo1 != 't' {
		goto _cgol_10
	}
	_nm_cgo2 = false
	size = int32(1)
	goto _cgol_6
_cgol_10:
	if _nm_cgo2 && _tag_cgo1 != 'L' {
		goto _cgol_11
	}
	_nm_cgo2 = false
	size = int32(2)
	goto _cgol_6
_cgol_11:
	if _nm_cgo2 && _tag_cgo1 != 'd' {
		goto _cgol_12
	}
	_nm_cgo2 = false
_cgol_12:
	if _nm_cgo2 && _tag_cgo1 != 'i' {
		goto _cgol_13
	}
	_nm_cgo2 = false
_cgol_13:
	if _nm_cgo2 && _tag_cgo1 != 'o' {
		goto _cgol_14
	}
	_nm_cgo2 = false
_cgol_14:
	if _nm_cgo2 && _tag_cgo1 != 'u' {
		goto _cgol_15
	}
	_nm_cgo2 = false
_cgol_15:
	if _nm_cgo2 && _tag_cgo1 != 'x' {
		goto _cgol_16
	}
	_nm_cgo2 = false
_cgol_16:
	if _nm_cgo2 && _tag_cgo1 != 'a' {
		goto _cgol_17
	}
	_nm_cgo2 = false
_cgol_17:
	if _nm_cgo2 && _tag_cgo1 != 'e' {
		goto _cgol_18
	}
	_nm_cgo2 = false
_cgol_18:
	if _nm_cgo2 && _tag_cgo1 != 'f' {
		goto _cgol_19
	}
	_nm_cgo2 = false
_cgol_19:
	if _nm_cgo2 && _tag_cgo1 != 'g' {
		goto _cgol_20
	}
	_nm_cgo2 = false
_cgol_20:
	if _nm_cgo2 && _tag_cgo1 != 'A' {
		goto _cgol_21
	}
	_nm_cgo2 = false
_cgol_21:
	if _nm_cgo2 && _tag_cgo1 != 'E' {
		goto _cgol_22
	}
	_nm_cgo2 = false
_cgol_22:
	if _nm_cgo2 && _tag_cgo1 != 'F' {
		goto _cgol_23
	}
	_nm_cgo2 = false
_cgol_23:
	if _nm_cgo2 && _tag_cgo1 != 'G' {
		goto _cgol_24
	}
	_nm_cgo2 = false
_cgol_24:
	if _nm_cgo2 && _tag_cgo1 != 'X' {
		goto _cgol_25
	}
	_nm_cgo2 = false
_cgol_25:
	if _nm_cgo2 && _tag_cgo1 != 's' {
		goto _cgol_26
	}
	_nm_cgo2 = false
_cgol_26:
	if _nm_cgo2 && _tag_cgo1 != 'c' {
		goto _cgol_27
	}
	_nm_cgo2 = false
_cgol_27:
	if _nm_cgo2 && _tag_cgo1 != '[' {
		goto _cgol_28
	}
	_nm_cgo2 = false
_cgol_28:
	if _nm_cgo2 && _tag_cgo1 != 'S' {
		goto _cgol_29
	}
	_nm_cgo2 = false
_cgol_29:
	if _nm_cgo2 && _tag_cgo1 != 'C' {
		goto _cgol_30
	}
	_nm_cgo2 = false
_cgol_30:
	if _nm_cgo2 && _tag_cgo1 != 'p' {
		goto _cgol_31
	}
	_nm_cgo2 = false
_cgol_31:
	if _nm_cgo2 && _tag_cgo1 != 'n' {
		goto _cgol_32
	}
	_nm_cgo2 = false
	*(*uintptr)(unsafe.Pointer(&p)) -= 4
	goto _cgol_6
_cgol_33:
	_nm_cgo2 = false
	goto fmt_fail
	goto _cgol_6
_cgol_32:
	goto _cgol_33
_cgol_6:
	t = *p
	if t&int32(47) == int32(3) {
		size = int32(1)
		t |= int32(32)
	}
	if t != 'n' {
		if t != '[' && t|int32(32) != 'c' {
			for iswspace(uint32(func() (_cgo_ret int32) {
				_cgo_addr := &c
				*_cgo_addr = int32(func() uint32 {
					if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) && int32(*f.Rpos) < int32(128) {
						return uint32(*func() (_cgo_ret *uint8) {
							_cgo_addr := &f.Rpos
							_cgo_ret = *_cgo_addr
							*(*uintptr)(unsafe.Pointer(_cgo_addr))++
							return
						}())
					} else {
						return getwc(f)
					}
				}())
				return *_cgo_addr
			}())) != 0 {
				pos++
			}
		} else {
			c = int32(func() uint32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) && int32(*f.Rpos) < int32(128) {
					return uint32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return getwc(f)
				}
			}())
		}
		if c < int32(0) {
			goto input_fail
		}
		func() uint32 {
			if f.Rend != nil && uint32(c) < uint32(128) {
				return uint32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return *_cgo_addr
				}())
			} else {
				return ungetwc(uint32(c), f)
			}
		}()
	}
	_tag_cgo3, _nm_cgo4 = t, true
	if _nm_cgo4 && _tag_cgo3 != 'n' {
		goto _cgol_34
	}
	_nm_cgo4 = false
	_cgos_store_int_vfwscanf(dest, size, uint64(pos))
	goto _cgol_1
_cgol_34:
	if _nm_cgo4 && _tag_cgo3 != 's' {
		goto _cgol_35
	}
	_nm_cgo4 = false
_cgol_35:
	if _nm_cgo4 && _tag_cgo3 != 'c' {
		goto _cgol_36
	}
	_nm_cgo4 = false
_cgol_36:
	if _nm_cgo4 && _tag_cgo3 != '[' {
		goto _cgol_37
	}
	_nm_cgo4 = false
	if !(t == 'c') {
		goto _cgol_39
	}
	if width < int32(1) {
		width = int32(1)
	}
	invert = int32(1)
	set = (*int32)(unsafe.Pointer(&[1]int32{'\x00'}))
	goto _cgol_38
_cgol_39:
	if !(t == 's') {
		goto _cgol_41
	}
	invert = int32(1)
	_cgos_vfwscanf_spaces_vfwscanf_cgo5 = [22]int32{' ', '\t', '\n', '\r', int32(11), int32(12), int32(133), int32(8192), int32(8193), int32(8194), int32(8195), int32(8196), int32(8197), int32(8198), int32(8200), int32(8201), int32(8202), int32(8232), int32(8233), int32(8287), int32(12288), int32(0)}
	set = (*int32)(unsafe.Pointer(&_cgos_vfwscanf_spaces_vfwscanf_cgo5))
	goto _cgol_40
_cgol_41:
	if *func() (_cgo_ret *int32) {
		_cgo_addr := &p
		*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
		return *_cgo_addr
	}() == '^' {
		func() int32 {
			*(*uintptr)(unsafe.Pointer(&p)) += 4
			return func() (_cgo_ret int32) {
				_cgo_addr := &invert
				*_cgo_addr = int32(1)
				return *_cgo_addr
			}()
		}()
	} else {
		invert = int32(0)
	}
	set = p
	if *p == ']' {
		*(*uintptr)(unsafe.Pointer(&p)) += 4
	}
_cgol_42:
	if !(*p != ']') {
		goto _cgol_43
	}
	if !!(*p != 0) {
		goto _cgol_44
	}
	goto fmt_fail
_cgol_44:
	*(*uintptr)(unsafe.Pointer(&p)) += 4
	goto _cgol_42
_cgol_43:
	;
_cgol_40:
	;
_cgol_38:
	s = (*int8)(func() unsafe.Pointer {
		if size == int32(0) {
			return dest
		} else {
			return nil
		}
	}())
	wcs = (*int32)(func() unsafe.Pointer {
		if size == int32(1) {
			return dest
		} else {
			return nil
		}
	}())
	gotmatch_cgo6 = int32(0)
	if width < int32(1) {
		width = -1
	}
	i = uint64(0)
	if alloc != 0 {
		k = uint64(func() uint32 {
			if t == 'c' {
				return uint32(width) + uint32(1)
			} else {
				return uint32(31)
			}
		}())
		if size == int32(1) {
			wcs = (*int32)(Malloc(k * 4))
			if !(wcs != nil) {
				goto alloc_fail
			}
		} else {
			s = (*int8)(Malloc(k))
			if !(s != nil) {
				goto alloc_fail
			}
		}
	}
	for width != 0 {
		if func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = int32(func() uint32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Rend)) && int32(*f.Rpos) < int32(128) {
					return uint32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return getwc(f)
				}
			}())
			return *_cgo_addr
		}() < int32(0) {
			break
		}
		if _cgos_in_set_vfwscanf(set, c) == invert {
			break
		}
		if wcs != nil {
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(wcs)) + uintptr(func() (_cgo_ret uint64) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}())*4)) = c
			if alloc != 0 && i == k {
				k += k + uint64(1)
				var tmp *int32 = (*int32)(Realloc(unsafe.Pointer(wcs), k*4))
				if !(tmp != nil) {
					goto alloc_fail
				}
				wcs = tmp
			}
		} else if size != int32(1) {
			var l int32 = Wctomb(func() *int8 {
				if s != nil {
					return (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))
				} else {
					return (*int8)(unsafe.Pointer(&tmp))
				}
			}(), c)
			if l < int32(0) {
				goto input_fail
			}
			i += uint64(l)
			if alloc != 0 && i > k-uint64(4) {
				k += k + uint64(1)
				var tmp *int8 = (*int8)(Realloc(unsafe.Pointer(s), k))
				if !(tmp != nil) {
					goto alloc_fail
				}
				s = tmp
			}
		}
		pos++
		width -= func() int32 {
			if width > int32(0) {
				return 1
			} else {
				return 0
			}
		}()
		gotmatch_cgo6 = int32(1)
	}
	if !(width != 0) {
		goto _cgol_45
	}
	func() uint32 {
		if f.Rend != nil && uint32(c) < uint32(128) {
			return uint32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return *_cgo_addr
			}())
		} else {
			return ungetwc(uint32(c), f)
		}
	}()
	if !(t == 'c' || !(gotmatch_cgo6 != 0)) {
		goto _cgol_46
	}
	goto match_fail
_cgol_46:
	;
_cgol_45:
	if alloc != 0 {
		if size == int32(1) {
			*(**int32)(dest) = wcs
		} else {
			*(**int8)(dest) = s
		}
	}
	if t != 'c' {
		if wcs != nil {
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(wcs)) + uintptr(i)*4)) = int32(0)
		}
		if s != nil {
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i))) = int8(0)
		}
	}
	goto _cgol_47
_cgol_37:
	if _nm_cgo4 && _tag_cgo3 != 'd' {
		goto _cgol_48
	}
	_nm_cgo4 = false
_cgol_48:
	if _nm_cgo4 && _tag_cgo3 != 'i' {
		goto _cgol_49
	}
	_nm_cgo4 = false
_cgol_49:
	if _nm_cgo4 && _tag_cgo3 != 'o' {
		goto _cgol_50
	}
	_nm_cgo4 = false
_cgol_50:
	if _nm_cgo4 && _tag_cgo3 != 'u' {
		goto _cgol_51
	}
	_nm_cgo4 = false
_cgol_51:
	if _nm_cgo4 && _tag_cgo3 != 'x' {
		goto _cgol_52
	}
	_nm_cgo4 = false
_cgol_52:
	if _nm_cgo4 && _tag_cgo3 != 'a' {
		goto _cgol_53
	}
	_nm_cgo4 = false
_cgol_53:
	if _nm_cgo4 && _tag_cgo3 != 'e' {
		goto _cgol_54
	}
	_nm_cgo4 = false
_cgol_54:
	if _nm_cgo4 && _tag_cgo3 != 'f' {
		goto _cgol_55
	}
	_nm_cgo4 = false
_cgol_55:
	if _nm_cgo4 && _tag_cgo3 != 'g' {
		goto _cgol_56
	}
	_nm_cgo4 = false
_cgol_56:
	if _nm_cgo4 && _tag_cgo3 != 'A' {
		goto _cgol_57
	}
	_nm_cgo4 = false
_cgol_57:
	if _nm_cgo4 && _tag_cgo3 != 'E' {
		goto _cgol_58
	}
	_nm_cgo4 = false
_cgol_58:
	if _nm_cgo4 && _tag_cgo3 != 'F' {
		goto _cgol_59
	}
	_nm_cgo4 = false
_cgol_59:
	if _nm_cgo4 && _tag_cgo3 != 'G' {
		goto _cgol_60
	}
	_nm_cgo4 = false
_cgol_60:
	if _nm_cgo4 && _tag_cgo3 != 'X' {
		goto _cgol_61
	}
	_nm_cgo4 = false
_cgol_61:
	if _nm_cgo4 && _tag_cgo3 != 'p' {
		goto _cgol_62
	}
	_nm_cgo4 = false
	if width < int32(1) {
		width = int32(0)
	}
	Snprintf((*int8)(unsafe.Pointer(&tmp)), 22, (*int8)(unsafe.Pointer(&[18]int8{'%', '.', '*', 's', '%', '.', '0', 'd', '%', 's', '%', 'c', '%', '%', 'l', 'l', 'n', '\x00'})), int32(1)+func() int32 {
		if !(dest != nil) {
			return 1
		} else {
			return 0
		}
	}(), (*int8)(unsafe.Pointer(&[3]int8{'%', '*', '\x00'})), width, (*int8)(unsafe.Pointer(&*(*[3]int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[3]int8)(unsafe.Pointer(&_cgos_vfwscanf_size_pfx_vfwscanf)))) + uintptr(size+int32(2))*3)))), t)
	cnt = int64(0)
	if !(Fscanf(f, (*int8)(unsafe.Pointer(&tmp)), func() unsafe.Pointer {
		if dest != nil {
			return dest
		} else {
			return unsafe.Pointer(&cnt)
		}
	}(), &cnt) == -1) {
		goto _cgol_64
	}
	goto input_fail
	goto _cgol_63
_cgol_64:
	if !!(cnt != 0) {
		goto _cgol_65
	}
	goto match_fail
_cgol_65:
	;
_cgol_63:
	pos += cnt
	goto _cgol_47
_cgol_66:
	_nm_cgo4 = false
	goto fmt_fail
	goto _cgol_47
_cgol_62:
	goto _cgol_66
_cgol_47:
	if dest != nil {
		matches++
	}
	*(*uintptr)(unsafe.Pointer(&p)) += 4
	goto _cgol_1
_cgol_2:
	if true {
		goto _cgol_67
	}
fmt_fail:
	;
alloc_fail:
	;
input_fail:
	if !(matches != 0) {
		matches--
	}
match_fail:
	if alloc != 0 {
		Free(unsafe.Pointer(s))
		Free(unsafe.Pointer(wcs))
	}
_cgol_67:
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return matches
}

var _cgos_vfwscanf_size_pfx_vfwscanf [6][3]int8 = [6][3]int8{[3]int8{'h', 'h', '\x00'}, [3]int8{'h', '\x00'}, [3]int8{'\x00'}, [3]int8{'l', '\x00'}, [3]int8{'L', '\x00'}, [3]int8{'l', 'l', '\x00'}}
