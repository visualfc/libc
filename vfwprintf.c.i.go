package libc

import unsafe "unsafe"

const (
	_cgoe_BARE_vfwprintf     int32 = 0
	_cgoe_LPRE_vfwprintf     int32 = 1
	_cgoe_LLPRE_vfwprintf    int32 = 2
	_cgoe_HPRE_vfwprintf     int32 = 3
	_cgoe_HHPRE_vfwprintf    int32 = 4
	_cgoe_BIGLPRE_vfwprintf  int32 = 5
	_cgoe_ZTPRE_vfwprintf    int32 = 6
	_cgoe_JPRE_vfwprintf     int32 = 7
	_cgoe_STOP_vfwprintf     int32 = 8
	_cgoe_PTR_vfwprintf      int32 = 9
	_cgoe_INT_vfwprintf      int32 = 10
	_cgoe_UINT_vfwprintf     int32 = 11
	_cgoe_ULLONG_vfwprintf   int32 = 12
	_cgoe_LONG_vfwprintf     int32 = 13
	_cgoe_ULONG_vfwprintf    int32 = 14
	_cgoe_SHORT_vfwprintf    int32 = 15
	_cgoe_USHORT_vfwprintf   int32 = 16
	_cgoe_CHAR_vfwprintf     int32 = 17
	_cgoe_UCHAR_vfwprintf    int32 = 18
	_cgoe_LLONG_vfwprintf    int32 = 19
	_cgoe_SIZET_vfwprintf    int32 = 20
	_cgoe_IMAX_vfwprintf     int32 = 21
	_cgoe_UMAX_vfwprintf     int32 = 22
	_cgoe_PDIFF_vfwprintf    int32 = 23
	_cgoe_UIPTR_vfwprintf    int32 = 24
	_cgoe_DBL_vfwprintf      int32 = 25
	_cgoe_LDBL_vfwprintf     int32 = 26
	_cgoe_NOARG_vfwprintf    int32 = 27
	_cgoe_MAXSTATE_vfwprintf int32 = 28
)

var _cgos_states_vfwprintf [8][58]uint8 = [8][58]uint8{[58]uint8{uint8(25), 0, uint8(10), 0, uint8(25), uint8(25), uint8(25), 0, 0, 0, 0, uint8(5), 0, 0, 0, 0, 0, 0, uint8(9), 0, 0, 0, 0, uint8(11), 0, 0, 0, 0, 0, 0, 0, 0, uint8(25), 0, uint8(17), uint8(10), uint8(25), uint8(25), uint8(25), uint8(3), uint8(10), uint8(7), 0, uint8(1), uint8(27), uint8(9), uint8(11), uint8(24), 0, 0, uint8(9), uint8(6), uint8(11), 0, 0, uint8(11), 0, uint8(6)}, [58]uint8{uint8(25), 0, 0, 0, uint8(25), uint8(25), uint8(25), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(14), 0, 0, 0, 0, 0, 0, 0, 0, uint8(25), 0, uint8(10), uint8(13), uint8(25), uint8(25), uint8(25), 0, uint8(13), 0, 0, uint8(2), 0, uint8(9), uint8(14), 0, 0, 0, uint8(9), 0, uint8(14), 0, 0, uint8(14)}, [58]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(12), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(19), 0, 0, 0, 0, uint8(19), 0, 0, 0, 0, uint8(9), uint8(12), 0, 0, 0, 0, 0, uint8(12), 0, 0, uint8(12)}, [58]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(16), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(15), 0, 0, 0, uint8(4), uint8(15), 0, 0, 0, 0, uint8(9), uint8(16), 0, 0, 0, 0, 0, uint8(16), 0, 0, uint8(16)}, [58]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(18), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(17), 0, 0, 0, 0, uint8(17), 0, 0, 0, 0, uint8(9), uint8(18), 0, 0, 0, 0, 0, uint8(18), 0, 0, uint8(18)}, [58]uint8{uint8(26), 0, 0, 0, uint8(26), uint8(26), uint8(26), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(26), 0, 0, 0, uint8(26), uint8(26), uint8(26), 0, 0, 0, 0, 0, 0, uint8(9)}, [58]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(20), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(23), 0, 0, 0, 0, uint8(23), 0, 0, 0, 0, uint8(9), uint8(20), 0, 0, 0, 0, 0, uint8(20), 0, 0, uint8(20)}, [58]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(22), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(21), 0, 0, 0, 0, uint8(21), 0, 0, 0, 0, uint8(9), uint8(22), 0, 0, 0, 0, 0, uint8(22), 0, 0, uint8(22)}}

type _cgos_arg_vfwprintf struct {
	i uint64
}

func _cgos_pop_arg_vfwprintf(arg *_cgos_arg_vfwprintf, type_ int32, ap *[]interface {
}) {
	switch type_ {
	case _cgoe_PTR_vfwprintf:
		*(*unsafe.Pointer)(unsafe.Pointer(arg)) = func(__cgo_args []interface {
		}) (_cgo_ret unsafe.Pointer) {
			_cgo_ret = unsafe.Pointer((*[2]unsafe.Pointer)(unsafe.Pointer(&__cgo_args[0]))[1])
			*ap = __cgo_args[1:]
			return
		}(*ap)
		break
	case _cgoe_INT_vfwprintf:
		arg.i = uint64(func(__cgo_args []interface {
		}) (_cgo_ret int32) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case int32:
				_cgo_ret = _cgo_tag
			case uint32:
				_cgo_ret = int32(_cgo_tag)
			case uint64:
				_cgo_ret = int32(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap))
		break
	case _cgoe_UINT_vfwprintf:
		arg.i = uint64(func(__cgo_args []interface {
		}) (_cgo_ret uint32) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case uint32:
				_cgo_ret = _cgo_tag
			case int32:
				_cgo_ret = uint32(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap))
		break
	case _cgoe_LONG_vfwprintf:
		arg.i = uint64(func(__cgo_args []interface {
		}) (_cgo_ret int64) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case int64:
				_cgo_ret = _cgo_tag
			case uint64:
				_cgo_ret = int64(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap))
		break
	case _cgoe_ULONG_vfwprintf:
		arg.i = uint64(func(__cgo_args []interface {
		}) (_cgo_ret uint64) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case uint64:
				_cgo_ret = _cgo_tag
			case int64:
				_cgo_ret = uint64(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap))
		break
	case _cgoe_ULLONG_vfwprintf:
		arg.i = func(__cgo_args []interface {
		}) (_cgo_ret uint64) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case uint64:
				_cgo_ret = _cgo_tag
			case int64:
				_cgo_ret = uint64(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap)
		break
	case _cgoe_SHORT_vfwprintf:
		arg.i = uint64(int16(func(__cgo_args []interface {
		}) (_cgo_ret int32) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case int32:
				_cgo_ret = _cgo_tag
			case uint32:
				_cgo_ret = int32(_cgo_tag)
			case uint64:
				_cgo_ret = int32(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap)))
		break
	case _cgoe_USHORT_vfwprintf:
		arg.i = uint64(uint16(func(__cgo_args []interface {
		}) (_cgo_ret int32) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case int32:
				_cgo_ret = _cgo_tag
			case uint32:
				_cgo_ret = int32(_cgo_tag)
			case uint64:
				_cgo_ret = int32(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap)))
		break
	case _cgoe_CHAR_vfwprintf:
		arg.i = uint64(int8(func(__cgo_args []interface {
		}) (_cgo_ret int32) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case int32:
				_cgo_ret = _cgo_tag
			case uint32:
				_cgo_ret = int32(_cgo_tag)
			case uint64:
				_cgo_ret = int32(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap)))
		break
	case _cgoe_UCHAR_vfwprintf:
		arg.i = uint64(uint8(func(__cgo_args []interface {
		}) (_cgo_ret int32) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case int32:
				_cgo_ret = _cgo_tag
			case uint32:
				_cgo_ret = int32(_cgo_tag)
			case uint64:
				_cgo_ret = int32(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap)))
		break
	case _cgoe_LLONG_vfwprintf:
		arg.i = uint64(func(__cgo_args []interface {
		}) (_cgo_ret int64) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case int64:
				_cgo_ret = _cgo_tag
			case uint64:
				_cgo_ret = int64(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap))
		break
	case _cgoe_SIZET_vfwprintf:
		arg.i = uint64(func(__cgo_args []interface {
		}) (_cgo_ret uint64) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case uint64:
				_cgo_ret = _cgo_tag
			case int64:
				_cgo_ret = uint64(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap))
		break
	case _cgoe_IMAX_vfwprintf:
		arg.i = uint64(func(__cgo_args []interface {
		}) (_cgo_ret int64) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case int64:
				_cgo_ret = _cgo_tag
			case uint64:
				_cgo_ret = int64(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap))
		break
	case _cgoe_UMAX_vfwprintf:
		arg.i = func(__cgo_args []interface {
		}) (_cgo_ret uint64) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case uint64:
				_cgo_ret = _cgo_tag
			case int64:
				_cgo_ret = uint64(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap)
		break
	case _cgoe_PDIFF_vfwprintf:
		arg.i = uint64(func(__cgo_args []interface {
		}) (_cgo_ret int64) {
			switch _cgo_tag := __cgo_args[0].(type) {
			case int64:
				_cgo_ret = _cgo_tag
			case uint64:
				_cgo_ret = int64(_cgo_tag)
			}
			*ap = __cgo_args[1:]
			return
		}(*ap))
		break
	case _cgoe_UIPTR_vfwprintf:
		arg.i = uint64(uint64(uintptr(func(__cgo_args []interface {
		}) (_cgo_ret unsafe.Pointer) {
			_cgo_ret = unsafe.Pointer((*[2]unsafe.Pointer)(unsafe.Pointer(&__cgo_args[0]))[1])
			*ap = __cgo_args[1:]
			return
		}(*ap))))
		break
	case _cgoe_DBL_vfwprintf:
		*(*float64)(unsafe.Pointer(arg)) = float64(func(__cgo_args []interface {
		}) (_cgo_ret float64) {
			_cgo_ret = __cgo_args[0].(float64)
			*ap = __cgo_args[1:]
			return
		}(*ap))
		break
	case _cgoe_LDBL_vfwprintf:
		*(*float64)(unsafe.Pointer(arg)) = func(__cgo_args []interface {
		}) (_cgo_ret float64) {
			_cgo_ret = __cgo_args[0].(float64)
			*ap = __cgo_args[1:]
			return
		}(*ap)
	}
}
func _cgos_out_vfwprintf(f *Struct__IO_FILE, s *int32, l uint64) {
	for func() (_cgo_ret uint64) {
		_cgo_addr := &l
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 && !(f.Flags&uint32(32) != 0) {
		fputwc(*func() (_cgo_ret *int32) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}(), f)
	}
}
func _cgos_getint_vfwprintf(s **int32) int32 {
	var i int32
	for i = int32(0); func() int32 {
		if int32(0) != 0 {
			return iswdigit(uint32(**s))
		} else {
			return func() int32 {
				if uint32(**s)-uint32('0') < uint32(10) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0; *(*uintptr)(unsafe.Pointer(&*s)) += 4 {
		if uint32(i) > 214748364 || **s-'0' > int32(2147483647)-int32(10)*i {
			i = -1
		} else {
			i = int32(10)*i + (**s - '0')
		}
	}
	return i
}

var _cgos_sizeprefix_vfwprintf [24]int8 = [24]int8{int8('L'), 0, 0, int8('j'), int8('L'), int8('L'), int8('L'), 0, int8('j'), 0, 0, 0, 0, 0, int8('j'), int8('j'), 0, 0, 0, 0, int8('j'), 0, 0, int8('j')}

func _cgos_wprintf_core_vfwprintf(f *Struct__IO_FILE, fmt *int32, ap *[]interface {
}, nl_arg *_cgos_arg_vfwprintf, nl_type *int32) int32 {
	var a *int32
	var z *int32
	var s *int32 = (*int32)(unsafe.Pointer(fmt))
	var l10n uint32 = uint32(0)
	var fl uint32
	var w int32
	var p int32
	var xp int32
	var arg _cgos_arg_vfwprintf
	var argpos int32
	var st uint32
	var ps uint32
	var cnt int32 = int32(0)
	var l int32 = int32(0)
	var i int32
	var t int32
	var bs *int8
	var charfmt [16]int8
	var wc int32
	for {
		if l > int32(2147483647)-cnt {
			goto overflow
		}
		cnt += l
		if !(*s != 0) {
			break
		}
		for a = s; *s != 0 && *s != '%'; *(*uintptr)(unsafe.Pointer(&s)) += 4 {
		}
		for z = s; *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(0))*4)) == '%' && *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))*4)) == '%'; func() *int32 {
			*(*uintptr)(unsafe.Pointer(&z)) += 4
			return func() (_cgo_ret *int32) {
				_cgo_addr := &s
				*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(int32(2)) * 4
				return *_cgo_addr
			}()
		}() {
		}
		if (uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(a)))/4 > uintptr(int64(int32(2147483647)-cnt)) {
			goto overflow
		}
		l = int32((uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(a))) / 4)
		if f != nil {
			_cgos_out_vfwprintf(f, a, uint64(l))
		}
		if l != 0 {
			continue
		}
		if func() int32 {
			if int32(0) != 0 {
				return iswdigit(uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))*4))))
			} else {
				return func() int32 {
					if uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))*4)))-uint32('0') < uint32(10) {
						return 1
					} else {
						return 0
					}
				}()
			}
		}() != 0 && *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2))*4)) == '$' {
			l10n = uint32(1)
			argpos = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))*4)) - '0'
			*(*uintptr)(unsafe.Pointer(&s)) += uintptr(int32(3)) * 4
		} else {
			argpos = -1
			*(*uintptr)(unsafe.Pointer(&s)) += 4
		}
		for fl = uint32(0); uint32(*s)-uint32(' ') < uint32(32) && 75913&(uint32(1)<<(*s-' ')) != 0; *(*uintptr)(unsafe.Pointer(&s)) += 4 {
			fl |= uint32(1) << (*s - ' ')
		}
		if *s == '*' {
			if func() int32 {
				if int32(0) != 0 {
					return iswdigit(uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))*4))))
				} else {
					return func() int32 {
						if uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))*4)))-uint32('0') < uint32(10) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}() != 0 && *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2))*4)) == '$' {
				l10n = uint32(1)
				*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))*4))-'0')*4)) = _cgoe_INT_vfwprintf
				w = int32((*(*_cgos_arg_vfwprintf)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_arg)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))*4))-'0')*8))).i)
				*(*uintptr)(unsafe.Pointer(&s)) += uintptr(int32(3)) * 4
			} else if !(l10n != 0) {
				w = func() int32 {
					if f != nil {
						return func(__cgo_args []interface {
						}) (_cgo_ret int32) {
							switch _cgo_tag := __cgo_args[0].(type) {
							case int32:
								_cgo_ret = _cgo_tag
							case uint32:
								_cgo_ret = int32(_cgo_tag)
							case uint64:
								_cgo_ret = int32(_cgo_tag)
							}
							*ap = __cgo_args[1:]
							return
						}(*ap)
					} else {
						return int32(0)
					}
				}()
				*(*uintptr)(unsafe.Pointer(&s)) += 4
			} else {
				goto inval
			}
			if w < int32(0) {
				func() int32 {
					fl |= 8192
					return func() (_cgo_ret int32) {
						_cgo_addr := &w
						*_cgo_addr = -w
						return *_cgo_addr
					}()
				}()
			}
		} else if func() (_cgo_ret int32) {
			_cgo_addr := &w
			*_cgo_addr = _cgos_getint_vfwprintf(&s)
			return *_cgo_addr
		}() < int32(0) {
			goto overflow
		}
		if *s == '.' && *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))*4)) == '*' {
			if func() int32 {
				if int32(0) != 0 {
					return Isdigit(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2))*4)))
				} else {
					return func() int32 {
						if uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2))*4)))-uint32('0') < uint32(10) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}() != 0 && *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(3))*4)) == '$' {
				*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2))*4))-'0')*4)) = _cgoe_INT_vfwprintf
				p = int32((*(*_cgos_arg_vfwprintf)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_arg)) + uintptr(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2))*4))-'0')*8))).i)
				*(*uintptr)(unsafe.Pointer(&s)) += uintptr(int32(4)) * 4
			} else if !(l10n != 0) {
				p = func() int32 {
					if f != nil {
						return func(__cgo_args []interface {
						}) (_cgo_ret int32) {
							switch _cgo_tag := __cgo_args[0].(type) {
							case int32:
								_cgo_ret = _cgo_tag
							case uint32:
								_cgo_ret = int32(_cgo_tag)
							case uint64:
								_cgo_ret = int32(_cgo_tag)
							}
							*ap = __cgo_args[1:]
							return
						}(*ap)
					} else {
						return int32(0)
					}
				}()
				*(*uintptr)(unsafe.Pointer(&s)) += uintptr(int32(2)) * 4
			} else {
				goto inval
			}
			xp = func() int32 {
				if p >= int32(0) {
					return 1
				} else {
					return 0
				}
			}()
		} else if *s == '.' {
			*(*uintptr)(unsafe.Pointer(&s)) += 4
			p = _cgos_getint_vfwprintf(&s)
			xp = int32(1)
		} else {
			p = -1
			xp = int32(0)
		}
		st = uint32(0)
		for {
			if uint32(*s)-uint32('A') > uint32(57) {
				goto inval
			}
			ps = st
			st = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[58]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[58]uint8)(unsafe.Pointer(&_cgos_states_vfwprintf)))) + uintptr(st)*58)))))) + uintptr(*func() (_cgo_ret *int32) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()-'A'))))
			if !(st-uint32(1) < uint32(8)) {
				break
			}
		}
		if !(st != 0) {
			goto inval
		}
		if st == uint32(27) {
			if argpos >= int32(0) {
				goto inval
			}
		} else if argpos >= int32(0) {
			func() _cgos_arg_vfwprintf {
				*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(argpos)*4)) = int32(st)
				return func() (_cgo_ret _cgos_arg_vfwprintf) {
					_cgo_addr := &arg
					*_cgo_addr = *(*_cgos_arg_vfwprintf)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_arg)) + uintptr(argpos)*8))
					return *_cgo_addr
				}()
			}()
		} else if f != nil {
			_cgos_pop_arg_vfwprintf(&arg, int32(st), ap)
		} else {
			return int32(0)
		}
		if !(f != nil) {
			continue
		}
		t = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) - uintptr(1)*4))
		if ps != 0 && t&int32(15) == int32(3) {
			t &= -33
		}
		switch t {
		case 'n':
			switch ps {
			case uint32(0):
				*(*int32)(*(*unsafe.Pointer)(unsafe.Pointer(&arg))) = cnt
				break
			case uint32(1):
				*(*int64)(*(*unsafe.Pointer)(unsafe.Pointer(&arg))) = int64(cnt)
				break
			case uint32(2):
				*(*int64)(*(*unsafe.Pointer)(unsafe.Pointer(&arg))) = int64(cnt)
				break
			case uint32(3):
				*(*uint16)(*(*unsafe.Pointer)(unsafe.Pointer(&arg))) = uint16(cnt)
				break
			case uint32(4):
				*(*uint8)(*(*unsafe.Pointer)(unsafe.Pointer(&arg))) = uint8(cnt)
				break
			case uint32(6):
				*(*uint64)(*(*unsafe.Pointer)(unsafe.Pointer(&arg))) = uint64(cnt)
				break
			case uint32(7):
				*(*uint64)(*(*unsafe.Pointer)(unsafe.Pointer(&arg))) = uint64(cnt)
				break
			}
			continue
		case 'c':
			if w < int32(1) {
				w = int32(1)
			}
			if w > int32(1) && !(fl&8192 != 0) {
				Fprintf(f, (*int8)(unsafe.Pointer(&[4]int8{'%', '*', 's', '\x00'})), w-int32(1), (*int8)(unsafe.Pointer(&[1]int8{'\x00'})))
			}
			fputwc(int32(btowc(int32(arg.i))), f)
			if w > int32(1) && fl&8192 != 0 {
				Fprintf(f, (*int8)(unsafe.Pointer(&[4]int8{'%', '*', 's', '\x00'})), w-int32(1), (*int8)(unsafe.Pointer(&[1]int8{'\x00'})))
			}
			l = w
			continue
		case 'C':
			fputwc(int32(arg.i), f)
			l = int32(1)
			continue
		case 'S':
			a = (*int32)(*(*unsafe.Pointer)(unsafe.Pointer(&arg)))
			z = (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(Wcsnlen(a, uint64(func() int32 {
				if p < int32(0) {
					return int32(2147483647)
				} else {
					return p
				}
			}())))*4))
			if p < int32(0) && *z != 0 {
				goto overflow
			}
			p = int32((uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(a))) / 4)
			if w < p {
				w = p
			}
			if !(fl&8192 != 0) {
				Fprintf(f, (*int8)(unsafe.Pointer(&[4]int8{'%', '*', 's', '\x00'})), w-p, (*int8)(unsafe.Pointer(&[1]int8{'\x00'})))
			}
			_cgos_out_vfwprintf(f, a, uint64(p))
			if fl&8192 != 0 {
				Fprintf(f, (*int8)(unsafe.Pointer(&[4]int8{'%', '*', 's', '\x00'})), w-p, (*int8)(unsafe.Pointer(&[1]int8{'\x00'})))
			}
			l = w
			continue
		case 'm':
			*(*unsafe.Pointer)(unsafe.Pointer(&arg)) = unsafe.Pointer(Strerror(*X__errno_location()))
		case 's':
			if !(*(*unsafe.Pointer)(unsafe.Pointer(&arg)) != nil) {
				*(*unsafe.Pointer)(unsafe.Pointer(&arg)) = unsafe.Pointer((*int8)(unsafe.Pointer(&[7]int8{'(', 'n', 'u', 'l', 'l', ')', '\x00'})))
			}
			bs = (*int8)(*(*unsafe.Pointer)(unsafe.Pointer(&arg)))
			for i = func() (_cgo_ret int32) {
				_cgo_addr := &l
				*_cgo_addr = int32(0)
				return *_cgo_addr
			}(); l < func() int32 {
				if p < int32(0) {
					return int32(2147483647)
				} else {
					return p
				}
			}() && func() (_cgo_ret int32) {
				_cgo_addr := &i
				*_cgo_addr = Mbtowc(&wc, bs, uint64(4))
				return *_cgo_addr
			}() > int32(0); func() int32 {
				*(*uintptr)(unsafe.Pointer(&bs)) += uintptr(i)
				return func() (_cgo_ret int32) {
					_cgo_addr := &l
					_cgo_ret = *_cgo_addr
					*_cgo_addr++
					return
				}()
			}() {
			}
			if i < int32(0) {
				return -1
			}
			if p < int32(0) && int32(*bs) != 0 {
				goto overflow
			}
			p = l
			if w < p {
				w = p
			}
			if !(fl&8192 != 0) {
				Fprintf(f, (*int8)(unsafe.Pointer(&[4]int8{'%', '*', 's', '\x00'})), w-p, (*int8)(unsafe.Pointer(&[1]int8{'\x00'})))
			}
			bs = (*int8)(*(*unsafe.Pointer)(unsafe.Pointer(&arg)))
			for func() (_cgo_ret int32) {
				_cgo_addr := &l
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				i = Mbtowc(&wc, bs, uint64(4))
				*(*uintptr)(unsafe.Pointer(&bs)) += uintptr(i)
				fputwc(wc, f)
			}
			if fl&8192 != 0 {
				Fprintf(f, (*int8)(unsafe.Pointer(&[4]int8{'%', '*', 's', '\x00'})), w-p, (*int8)(unsafe.Pointer(&[1]int8{'\x00'})))
			}
			l = w
			continue
		}
		if xp != 0 && p < int32(0) {
			goto overflow
		}
		Snprintf((*int8)(unsafe.Pointer(&charfmt)), 16, (*int8)(unsafe.Pointer(&[20]int8{'%', '%', '%', 's', '%', 's', '%', 's', '%', 's', '%', 's', '*', '.', '*', '%', 'c', '%', 'c', '\x00'})), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[2]int8{'#', '\x00'}))))+func() uintptr {
			if !(fl&8 != 0) {
				return 1
			} else {
				return 0
			}
		}())), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[2]int8{'+', '\x00'}))))+func() uintptr {
			if !(fl&2048 != 0) {
				return 1
			} else {
				return 0
			}
		}())), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[2]int8{'-', '\x00'}))))+func() uintptr {
			if !(fl&8192 != 0) {
				return 1
			} else {
				return 0
			}
		}())), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[2]int8{' ', '\x00'}))))+func() uintptr {
			if !(fl&1 != 0) {
				return 1
			} else {
				return 0
			}
		}())), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[2]int8{'0', '\x00'}))))+func() uintptr {
			if !(fl&65536 != 0) {
				return 1
			} else {
				return 0
			}
		}())), int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_sizeprefix_vfwprintf)))) + uintptr(t|int32(32)-'a')))), t)
		switch t | int32(32) {
		case 'a':
			fallthrough
		case 'e':
			fallthrough
		case 'f':
			fallthrough
		case 'g':
			l = Fprintf(f, (*int8)(unsafe.Pointer(&charfmt)), w, p, *(*float64)(unsafe.Pointer(&arg)))
			break
		case 'd':
			fallthrough
		case 'i':
			fallthrough
		case 'o':
			fallthrough
		case 'u':
			fallthrough
		case 'x':
			fallthrough
		case 'p':
			l = Fprintf(f, (*int8)(unsafe.Pointer(&charfmt)), w, p, arg.i)
			break
		}
	}
	if f != nil {
		return cnt
	}
	if !(l10n != 0) {
		return int32(0)
	}
	for i = int32(1); i <= int32(9) && *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(i)*4)) != 0; i++ {
		_cgos_pop_arg_vfwprintf((*_cgos_arg_vfwprintf)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_arg))+uintptr(i)*8)), *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(i)*4)), ap)
	}
	for ; i <= int32(9) && !(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(i)*4)) != 0); i++ {
	}
	if i <= int32(9) {
		return -1
	}
	return int32(1)
inval:
	*X__errno_location() = int32(22)
	return -1
overflow:
	*X__errno_location() = int32(75)
	return -1
}
func vfwprintf(f *Struct__IO_FILE, fmt *int32, ap []interface {
}) int32 {
	var ap2 []interface {
	}
	var nl_type [9]int32 = [9]int32{int32(0)}
	var nl_arg [9]_cgos_arg_vfwprintf
	var olderr int32
	var ret int32
	ap2 = ap
	if _cgos_wprintf_core_vfwprintf(nil, fmt, &ap2, (*_cgos_arg_vfwprintf)(unsafe.Pointer(&nl_arg)), (*int32)(unsafe.Pointer(&nl_type))) < int32(0) {
		return -1
	}
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	fwide(f, int32(1))
	olderr = int32(f.Flags & uint32(32))
	f.Flags &= uint32(4294967263)
	ret = _cgos_wprintf_core_vfwprintf(f, fmt, &ap2, (*_cgos_arg_vfwprintf)(unsafe.Pointer(&nl_arg)), (*int32)(unsafe.Pointer(&nl_type)))
	if f.Flags&uint32(32) != 0 {
		ret = -1
	}
	f.Flags |= uint32(olderr)
	for {
		if __need_unlock != 0 {
			__unlockfile(f)
		}
		if true {
			break
		}
	}
	return ret
}
