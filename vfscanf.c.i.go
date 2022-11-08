package libc

import unsafe "unsafe"

func _cgos_store_int_vfscanf(dest unsafe.Pointer, size int32, i uint64) {
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
func _cgos_arg_n_vfscanf(ap []interface {
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
func Vfscanf(f *Struct__IO_FILE, fmt *int8, ap []interface {
}) int32 {
	var (
		_tag_cgo1 int32
		_nm_cgo2  bool
		_tag_cgo3 int32
		_nm_cgo4  bool
	)
	var width int32
	var size int32
	var alloc int32 = int32(0)
	var base int32
	var p *uint8
	var c int32
	var t int32
	var s *int8
	var wcs *int32
	var st Struct___mbstate_t
	var dest unsafe.Pointer = unsafe.Pointer(nil)
	var invert int32
	var matches int32 = int32(0)
	var x uint64
	var y float64
	var pos int64 = int64(0)
	var scanset [257]uint8
	var i uint64
	var k uint64
	var wc int32
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	if !(f.Rpos != nil) {
		__toread(f)
	}
	if !(f.Rpos != nil) {
		goto input_fail
	}
	p = (*uint8)(unsafe.Pointer(fmt))
_cgol_1:
	if !(*p != 0) {
		goto _cgol_2
	}
	alloc = int32(0)
	if X__isspace(int32(*p)) != 0 {
		for X__isspace(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1)))))) != 0 {
			*(*uintptr)(unsafe.Pointer(&p))++
		}
		__shlim(f, int64(0))
		for X__isspace(func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}()) != 0 {
		}
		if f.Shlim >= int64(0) {
			func() int {
				_ = func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return
				}()
				return 0
			}()
		} else {
			func() int {
				_ = int32(0)
				return 0
			}()
		}
		pos += f.Shcnt + int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Buf)))
		goto _cgol_1
	}
	if !(int32(*p) != '%' || int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))))) == '%') {
		goto _cgol_3
	}
	__shlim(f, int64(0))
	if int32(*p) == '%' {
		*(*uintptr)(unsafe.Pointer(&p))++
		for X__isspace(func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = func() int32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
					return int32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return __shgetc(f)
				}
			}()
			return *_cgo_addr
		}()) != 0 {
		}
	} else {
		c = func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}()
	}
	if !(c != int32(*p)) {
		goto _cgol_4
	}
	if f.Shlim >= int64(0) {
		func() int {
			_ = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return
			}()
			return 0
		}()
	} else {
		func() int {
			_ = int32(0)
			return 0
		}()
	}
	if c < int32(0) {
		goto input_fail
	}
	goto match_fail
_cgol_4:
	pos += f.Shcnt + int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Buf)))
	goto _cgol_1
_cgol_3:
	*(*uintptr)(unsafe.Pointer(&p))++
	if int32(*p) == '*' {
		dest = unsafe.Pointer(nil)
		*(*uintptr)(unsafe.Pointer(&p))++
	} else if func() int32 {
		if int32(0) != 0 {
			return Isdigit(int32(*p))
		} else {
			return func() int32 {
				if uint32(*p)-uint32('0') < uint32(10) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))))) == '$' {
		dest = _cgos_arg_n_vfscanf(ap, uint32(int32(*p)-'0'))
		*(*uintptr)(unsafe.Pointer(&p)) += uintptr(int32(2))
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
			return Isdigit(int32(*p))
		} else {
			return func() int32 {
				if uint32(*p)-uint32('0') < uint32(10) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0; *(*uintptr)(unsafe.Pointer(&p))++ {
		width = int32(10)*width + int32(*p) - '0'
	}
	if int32(*p) == 'm' {
		wcs = (*int32)(nil)
		s = (*int8)(nil)
		alloc = func() int32 {
			if !!(dest != nil) {
				return 1
			} else {
				return 0
			}
		}()
		*(*uintptr)(unsafe.Pointer(&p))++
	} else {
		alloc = int32(0)
	}
	size = int32(0)
	_tag_cgo1, _nm_cgo2 = int32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}()), true
	if _nm_cgo2 && _tag_cgo1 != 'h' {
		goto _cgol_5
	}
	_nm_cgo2 = false
	if int32(*p) == 'h' {
		func() int32 {
			*(*uintptr)(unsafe.Pointer(&p))++
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
	if int32(*p) == 'l' {
		func() int32 {
			*(*uintptr)(unsafe.Pointer(&p))++
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
	*(*uintptr)(unsafe.Pointer(&p))--
	goto _cgol_6
_cgol_33:
	_nm_cgo2 = false
	goto fmt_fail
	goto _cgol_6
_cgol_32:
	goto _cgol_33
_cgol_6:
	t = int32(*p)
	if t&int32(47) == int32(3) {
		t |= int32(32)
		size = int32(1)
	}
	switch t {
	case 'c':
		if width < int32(1) {
			width = int32(1)
		}
	case '[':
		break
	case 'n':
		_cgos_store_int_vfscanf(dest, size, uint64(pos))
		goto _cgol_1
	default:
		__shlim(f, int64(0))
		for X__isspace(func() int32 {
			if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
				return int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}())
			} else {
				return __shgetc(f)
			}
		}()) != 0 {
		}
		if f.Shlim >= int64(0) {
			func() int {
				_ = func() (_cgo_ret *uint8) {
					_cgo_addr := &f.Rpos
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return
				}()
				return 0
			}()
		} else {
			func() int {
				_ = int32(0)
				return 0
			}()
		}
		pos += f.Shcnt + int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Buf)))
	}
	__shlim(f, int64(width))
	if func() int32 {
		if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
			return int32(*func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())
		} else {
			return __shgetc(f)
		}
	}() < int32(0) {
		goto input_fail
	}
	if f.Shlim >= int64(0) {
		func() int {
			_ = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return
			}()
			return 0
		}()
	} else {
		func() int {
			_ = int32(0)
			return 0
		}()
	}
	_tag_cgo3, _nm_cgo4 = t, true
	if _nm_cgo4 && _tag_cgo3 != 's' {
		goto _cgol_34
	}
	_nm_cgo4 = false
_cgol_34:
	if _nm_cgo4 && _tag_cgo3 != 'c' {
		goto _cgol_35
	}
	_nm_cgo4 = false
_cgol_35:
	if _nm_cgo4 && _tag_cgo3 != '[' {
		goto _cgol_36
	}
	_nm_cgo4 = false
	if !(t == 'c' || t == 's') {
		goto _cgol_38
	}
	Memset(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset))), -1, 257)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(int32(0)))) = uint8(0)
	if t == 's' {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(10))) = uint8(0)
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(11))) = uint8(0)
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(12))) = uint8(0)
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(13))) = uint8(0)
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(14))) = uint8(0)
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(33))) = uint8(0)
	}
	goto _cgol_37
_cgol_38:
	if int32(*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return *_cgo_addr
	}()) == '^' {
		func() int32 {
			*(*uintptr)(unsafe.Pointer(&p))++
			return func() (_cgo_ret int32) {
				_cgo_addr := &invert
				*_cgo_addr = int32(1)
				return *_cgo_addr
			}()
		}()
	} else {
		invert = int32(0)
	}
	Memset(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset))), invert, 257)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(int32(0)))) = uint8(0)
	if int32(*p) == '-' {
		func() uint8 {
			*(*uintptr)(unsafe.Pointer(&p))++
			return func() (_cgo_ret uint8) {
				_cgo_addr := &*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(46)))
				*_cgo_addr = uint8(int32(1) - invert)
				return *_cgo_addr
			}()
		}()
	} else if int32(*p) == ']' {
		func() uint8 {
			*(*uintptr)(unsafe.Pointer(&p))++
			return func() (_cgo_ret uint8) {
				_cgo_addr := &*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(94)))
				*_cgo_addr = uint8(int32(1) - invert)
				return *_cgo_addr
			}()
		}()
	}
_cgol_39:
	if !(int32(*p) != ']') {
		goto _cgol_40
	}
	if !!(*p != 0) {
		goto _cgol_41
	}
	goto fmt_fail
_cgol_41:
	if int32(*p) == '-' && int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))))) != 0 && int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))))) != ']' {
		for c = int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(func() (_cgo_ret *uint8) {
			_cgo_addr := &p
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}())) - uintptr(1)))); c < int32(*p); c++ {
			*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(int32(1)+c))) = uint8(int32(1) - invert)
		}
	}
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(int32(1)+int32(*p)))) = uint8(int32(1) - invert)
	*(*uintptr)(unsafe.Pointer(&p))++
	goto _cgol_39
_cgol_40:
	;
_cgol_37:
	wcs = (*int32)(nil)
	s = (*int8)(nil)
	i = uint64(0)
	k = uint64(func() uint32 {
		if t == 'c' {
			return uint32(width) + uint32(1)
		} else {
			return uint32(31)
		}
	}())
	if size == int32(1) {
		if alloc != 0 {
			wcs = (*int32)(Malloc(k * 4))
			if !(wcs != nil) {
				goto alloc_fail
			}
		} else {
			wcs = (*int32)(dest)
		}
		st = Struct___mbstate_t{uint32(0), 0}
		for *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = func() int32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
					return int32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return __shgetc(f)
				}
			}()
			return *_cgo_addr
		}()+int32(1)))) != 0 {
			switch mbrtowc(&wc, func() *int8 {
				_c := int8(c)
				return &_c
			}(), uint64(1), &st) {
			case uint64(18446744073709551615):
				goto input_fail
			case uint64(18446744073709551614):
				continue
			}
			if wcs != nil {
				*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(wcs)) + uintptr(func() (_cgo_ret uint64) {
					_cgo_addr := &i
					_cgo_ret = *_cgo_addr
					*_cgo_addr++
					return
				}())*4)) = wc
			}
			if alloc != 0 && i == k {
				k += k + uint64(1)
				var tmp *int32 = (*int32)(Realloc(unsafe.Pointer(wcs), k*4))
				if !(tmp != nil) {
					goto alloc_fail
				}
				wcs = tmp
			}
		}
		if !(mbsinit(&st) != 0) {
			goto input_fail
		}
	} else if alloc != 0 {
		s = (*int8)(Malloc(k))
		if !(s != nil) {
			goto alloc_fail
		}
		for *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = func() int32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
					return int32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return __shgetc(f)
				}
			}()
			return *_cgo_addr
		}()+int32(1)))) != 0 {
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(func() (_cgo_ret uint64) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}()))) = int8(c)
			if i == k {
				k += k + uint64(1)
				var tmp *int8 = (*int8)(Realloc(unsafe.Pointer(s), k))
				if !(tmp != nil) {
					goto alloc_fail
				}
				s = tmp
			}
		}
	} else if func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = (*int8)(dest)
		return *_cgo_addr
	}() != nil {
		for *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = func() int32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
					return int32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return __shgetc(f)
				}
			}()
			return *_cgo_addr
		}()+int32(1)))) != 0 {
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(func() (_cgo_ret uint64) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}()))) = int8(c)
		}
	} else {
		for *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&scanset)))) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &c
			*_cgo_addr = func() int32 {
				if uintptr(unsafe.Pointer(f.Rpos)) != uintptr(unsafe.Pointer(f.Shend)) {
					return int32(*func() (_cgo_ret *uint8) {
						_cgo_addr := &f.Rpos
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr))++
						return
					}())
				} else {
					return __shgetc(f)
				}
			}()
			return *_cgo_addr
		}()+int32(1)))) != 0 {
		}
	}
	if f.Shlim >= int64(0) {
		func() int {
			_ = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Rpos
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return
			}()
			return 0
		}()
	} else {
		func() int {
			_ = int32(0)
			return 0
		}()
	}
	if !!(f.Shcnt+int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Buf))) != 0) {
		goto _cgol_42
	}
	goto match_fail
_cgol_42:
	if !(t == 'c' && f.Shcnt+int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Buf))) != int64(width)) {
		goto _cgol_43
	}
	goto match_fail
_cgol_43:
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
	goto _cgol_44
_cgol_36:
	if _nm_cgo4 && _tag_cgo3 != 'p' {
		goto _cgol_45
	}
	_nm_cgo4 = false
_cgol_45:
	if _nm_cgo4 && _tag_cgo3 != 'X' {
		goto _cgol_46
	}
	_nm_cgo4 = false
_cgol_46:
	if _nm_cgo4 && _tag_cgo3 != 'x' {
		goto _cgol_47
	}
	_nm_cgo4 = false
	base = int32(16)
	goto int_common
_cgol_47:
	if _nm_cgo4 && _tag_cgo3 != 'o' {
		goto _cgol_48
	}
	_nm_cgo4 = false
	base = int32(8)
	goto int_common
_cgol_48:
	if _nm_cgo4 && _tag_cgo3 != 'd' {
		goto _cgol_49
	}
	_nm_cgo4 = false
_cgol_49:
	if _nm_cgo4 && _tag_cgo3 != 'u' {
		goto _cgol_50
	}
	_nm_cgo4 = false
	base = int32(10)
	goto int_common
_cgol_50:
	if _nm_cgo4 && _tag_cgo3 != 'i' {
		goto _cgol_51
	}
	_nm_cgo4 = false
	base = int32(0)
int_common:
	x = __intscan(f, uint32(base), int32(0), 18446744073709551615)
	if !!(f.Shcnt+int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Buf))) != 0) {
		goto _cgol_52
	}
	goto match_fail
_cgol_52:
	if t == 'p' && dest != nil {
		*(*unsafe.Pointer)(dest) = unsafe.Pointer(uintptr(uint64(x)))
	} else {
		_cgos_store_int_vfscanf(dest, size, x)
	}
	goto _cgol_44
_cgol_51:
	if _nm_cgo4 && _tag_cgo3 != 'a' {
		goto _cgol_53
	}
	_nm_cgo4 = false
_cgol_53:
	if _nm_cgo4 && _tag_cgo3 != 'A' {
		goto _cgol_54
	}
	_nm_cgo4 = false
_cgol_54:
	if _nm_cgo4 && _tag_cgo3 != 'e' {
		goto _cgol_55
	}
	_nm_cgo4 = false
_cgol_55:
	if _nm_cgo4 && _tag_cgo3 != 'E' {
		goto _cgol_56
	}
	_nm_cgo4 = false
_cgol_56:
	if _nm_cgo4 && _tag_cgo3 != 'f' {
		goto _cgol_57
	}
	_nm_cgo4 = false
_cgol_57:
	if _nm_cgo4 && _tag_cgo3 != 'F' {
		goto _cgol_58
	}
	_nm_cgo4 = false
_cgol_58:
	if _nm_cgo4 && _tag_cgo3 != 'g' {
		goto _cgol_59
	}
	_nm_cgo4 = false
_cgol_59:
	if _nm_cgo4 && _tag_cgo3 != 'G' {
		goto _cgol_60
	}
	_nm_cgo4 = false
	y = __floatscan(f, size, int32(0))
	if !!(f.Shcnt+int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Buf))) != 0) {
		goto _cgol_61
	}
	goto match_fail
_cgol_61:
	if dest != nil {
		switch size {
		case int32(0):
			*(*float32)(dest) = float32(y)
			break
		case int32(1):
			*(*float64)(dest) = float64(y)
			break
		case int32(2):
			*(*float64)(dest) = y
			break
		}
	}
	goto _cgol_44
	goto _cgol_44
_cgol_60:
	;
_cgol_44:
	pos += f.Shcnt + int64(uintptr(unsafe.Pointer(f.Rpos))-uintptr(unsafe.Pointer(f.Buf)))
	if dest != nil {
		matches++
	}
	*(*uintptr)(unsafe.Pointer(&p))++
	goto _cgol_1
_cgol_2:
	if true {
		goto _cgol_62
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
_cgol_62:
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
