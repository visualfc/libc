package libc

import unsafe "unsafe"

const (
	BARE     int32 = 0
	LPRE     int32 = 1
	LLPRE    int32 = 2
	HPRE     int32 = 3
	HHPRE    int32 = 4
	BIGLPRE  int32 = 5
	ZTPRE    int32 = 6
	JPRE     int32 = 7
	STOP     int32 = 8
	PTR      int32 = 9
	INT      int32 = 10
	UINT     int32 = 11
	ULLONG   int32 = 12
	LONG     int32 = 13
	ULONG    int32 = 14
	SHORT    int32 = 15
	USHORT   int32 = 16
	CHAR     int32 = 17
	UCHAR    int32 = 18
	LLONG    int32 = 19
	SIZET    int32 = 20
	IMAX     int32 = 21
	UMAX     int32 = 22
	PDIFF    int32 = 23
	UIPTR    int32 = 24
	DBL      int32 = 25
	LDBL     int32 = 26
	NOARG    int32 = 27
	MAXSTATE int32 = 28
)

var _cgos_states_vfprintf [8][58]uint8 = [8][58]uint8{[58]uint8{uint8(25), 0, uint8(10), 0, uint8(25), uint8(25), uint8(25), 0, 0, 0, 0, uint8(5), 0, 0, 0, 0, 0, 0, uint8(9), 0, 0, 0, 0, uint8(11), 0, 0, 0, 0, 0, 0, 0, 0, uint8(25), 0, uint8(17), uint8(10), uint8(25), uint8(25), uint8(25), uint8(3), uint8(10), uint8(7), 0, uint8(1), uint8(27), uint8(9), uint8(11), uint8(24), 0, 0, uint8(9), uint8(6), uint8(11), 0, 0, uint8(11), 0, uint8(6)}, [58]uint8{uint8(25), 0, 0, 0, uint8(25), uint8(25), uint8(25), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(14), 0, 0, 0, 0, 0, 0, 0, 0, uint8(25), 0, uint8(10), uint8(13), uint8(25), uint8(25), uint8(25), 0, uint8(13), 0, 0, uint8(2), 0, uint8(9), uint8(14), 0, 0, 0, uint8(9), 0, uint8(14), 0, 0, uint8(14)}, [58]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(12), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(19), 0, 0, 0, 0, uint8(19), 0, 0, 0, 0, uint8(9), uint8(12), 0, 0, 0, 0, 0, uint8(12), 0, 0, uint8(12)}, [58]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(16), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(15), 0, 0, 0, uint8(4), uint8(15), 0, 0, 0, 0, uint8(9), uint8(16), 0, 0, 0, 0, 0, uint8(16), 0, 0, uint8(16)}, [58]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(18), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(17), 0, 0, 0, 0, uint8(17), 0, 0, 0, 0, uint8(9), uint8(18), 0, 0, 0, 0, 0, uint8(18), 0, 0, uint8(18)}, [58]uint8{uint8(26), 0, 0, 0, uint8(26), uint8(26), uint8(26), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(26), 0, 0, 0, uint8(26), uint8(26), uint8(26), 0, 0, 0, 0, 0, 0, uint8(9)}, [58]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(20), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(23), 0, 0, 0, 0, uint8(23), 0, 0, 0, 0, uint8(9), uint8(20), 0, 0, 0, 0, 0, uint8(20), 0, 0, uint8(20)}, [58]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(22), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint8(21), 0, 0, 0, 0, uint8(21), 0, 0, 0, 0, uint8(9), uint8(22), 0, 0, 0, 0, 0, uint8(22), 0, 0, uint8(22)}}

type _cgos_arg_vfprintf struct {
	i uint64
}

func _cgos_pop_arg_vfprintf(arg *_cgos_arg_vfprintf, type_ int32, ap *[]interface {
}) {
	switch type_ {
	case PTR:
		*(*unsafe.Pointer)(unsafe.Pointer(arg)) = func(__cgo_args []interface {
		}) (_cgo_ret unsafe.Pointer) {
			_cgo_ret = unsafe.Pointer((*[2]unsafe.Pointer)(unsafe.Pointer(&__cgo_args[0]))[1])
			*ap = __cgo_args[1:]
			return
		}(*ap)
		break
	case INT:
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
	case UINT:
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
	case LONG:
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
	case ULONG:
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
	case ULLONG:
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
	case SHORT:
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
	case USHORT:
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
	case CHAR:
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
	case UCHAR:
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
	case LLONG:
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
	case SIZET:
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
	case IMAX:
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
	case UMAX:
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
	case PDIFF:
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
	case UIPTR:
		arg.i = uint64(uint64(uintptr(func(__cgo_args []interface {
		}) (_cgo_ret unsafe.Pointer) {
			_cgo_ret = unsafe.Pointer((*[2]unsafe.Pointer)(unsafe.Pointer(&__cgo_args[0]))[1])
			*ap = __cgo_args[1:]
			return
		}(*ap))))
		break
	case DBL:
		*(*float64)(unsafe.Pointer(arg)) = float64(func(__cgo_args []interface {
		}) (_cgo_ret float64) {
			_cgo_ret = __cgo_args[0].(float64)
			*ap = __cgo_args[1:]
			return
		}(*ap))
		break
	case LDBL:
		*(*float64)(unsafe.Pointer(arg)) = func(__cgo_args []interface {
		}) (_cgo_ret float64) {
			_cgo_ret = __cgo_args[0].(float64)
			*ap = __cgo_args[1:]
			return
		}(*ap)
	}
}
func _cgos_out_vfprintf(f *Struct__IO_FILE, s *int8, l uint64) {
	if !(f.Flags&uint32(32) != 0) {
		__fwritex((*uint8)(unsafe.Pointer(s)), l, f)
	}
}
func _cgos_pad_vfprintf(f *Struct__IO_FILE, c int8, w int32, l int32, fl int32) {
	var pad [256]int8
	if uint32(fl)&73728 != 0 || l >= w {
		return
	}
	l = w - l
	Memset(unsafe.Pointer((*int8)(unsafe.Pointer(&pad))), int32(c), func() uint64 {
		if uint64(l) > 256 {
			return 256
		} else {
			return uint64(l)
		}
	}())
	for ; uint64(l) >= 256; l -= int32(256) {
		_cgos_out_vfprintf(f, (*int8)(unsafe.Pointer(&pad)), 256)
	}
	_cgos_out_vfprintf(f, (*int8)(unsafe.Pointer(&pad)), uint64(l))
}

var _cgos_xdigits_vfprintf [16]int8 = [16]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}

func _cgos_fmt_x_vfprintf(x uint64, s *int8, lower int32) *int8 {
	for ; x != 0; x >>= int32(4) {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = int8(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_xdigits_vfprintf)))) + uintptr(x&uint64(15))))) | lower)
	}
	return s
}
func _cgos_fmt_o_vfprintf(x uint64, s *int8) *int8 {
	for ; x != 0; x >>= int32(3) {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = int8(uint64('0') + x&uint64(7))
	}
	return s
}
func _cgos_fmt_u_vfprintf(x uint64, s *int8) *int8 {
	var y uint64
	for ; x > uint64(4294967295); x /= uint64(10) {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = int8(uint64('0') + x%uint64(10))
	}
	for y = uint64(x); y != 0; y /= uint64(10) {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = int8(uint64('0') + y%uint64(10))
	}
	return s
}

type compiler_defines_long_double_incorrectly = int32

func _cgos_fmt_fp_vfprintf(f *Struct__IO_FILE, y float64, w int32, p int32, fl int32, t int32) int32 {
	var big [126]uint32
	var a *uint32
	var d *uint32
	var r *uint32
	var z *uint32
	var e2 int32 = int32(0)
	var e int32
	var i int32
	var j int32
	var l int32
	var buf [22]int8
	var s *int8
	var prefix *int8 = (*int8)(unsafe.Pointer(&[19]int8{'-', '0', 'X', '+', '0', 'X', ' ', '0', 'X', '-', '0', 'x', '+', '0', 'x', ' ', '0', 'x', '\x00'}))
	var pl int32
	var ebuf0 [12]int8
	var ebuf *int8 = &*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&ebuf0)))) + uintptr(12)))
	var estr *int8
	pl = int32(1)
	if func() int32 {
		if 8 == 4 {
			return int32(X__FLOAT_BITS(float32(y)) >> int32(31))
		} else {
			return func() int32 {
				if 8 == 8 {
					return int32(X__DOUBLE_BITS(float64(y)) >> int32(63))
				} else {
					return X__signbitl(y)
				}
			}()
		}
	}() != 0 {
		y = -y
	} else if uint32(fl)&2048 != 0 {
		*(*uintptr)(unsafe.Pointer(&prefix)) += uintptr(int32(3))
	} else if uint32(fl)&1 != 0 {
		*(*uintptr)(unsafe.Pointer(&prefix)) += uintptr(int32(6))
	} else {
		func() int32 {
			*(*uintptr)(unsafe.Pointer(&prefix))++
			return func() (_cgo_ret int32) {
				_cgo_addr := &pl
				*_cgo_addr = int32(0)
				return *_cgo_addr
			}()
		}()
	}
	if !(func() int32 {
		if 8 == 4 {
			return func() int32 {
				if X__FLOAT_BITS(float32(y))&uint32(2147483647) < uint32(2139095040) {
					return 1
				} else {
					return 0
				}
			}()
		} else {
			return func() int32 {
				if 8 == 8 {
					return func() int32 {
						if X__DOUBLE_BITS(float64(y))&9223372036854775807 < 9218868437227405312 {
							return 1
						} else {
							return 0
						}
					}()
				} else {
					return func() int32 {
						if X__fpclassifyl(y) > int32(1) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}()
		}
	}() != 0) {
		var s *int8 = func() *int8 {
			if t&int32(32) != 0 {
				return (*int8)(unsafe.Pointer(&[4]int8{'i', 'n', 'f', '\x00'}))
			} else {
				return (*int8)(unsafe.Pointer(&[4]int8{'I', 'N', 'F', '\x00'}))
			}
		}()
		if y != y {
			s = func() *int8 {
				if t&int32(32) != 0 {
					return (*int8)(unsafe.Pointer(&[4]int8{'n', 'a', 'n', '\x00'}))
				} else {
					return (*int8)(unsafe.Pointer(&[4]int8{'N', 'A', 'N', '\x00'}))
				}
			}()
		}
		_cgos_pad_vfprintf(f, int8(' '), w, int32(3)+pl, int32(uint32(fl)&4294901759))
		_cgos_out_vfprintf(f, prefix, uint64(pl))
		_cgos_out_vfprintf(f, s, uint64(3))
		_cgos_pad_vfprintf(f, int8(' '), w, int32(3)+pl, int32(uint32(fl)^8192))
		return func() int32 {
			if w > int32(3)+pl {
				return w
			} else {
				return int32(3) + pl
			}
		}()
	}
	y = Frexpl(y, &e2) * float64(int32(2))
	if y != 0 {
		e2--
	}
	if t|int32(32) == 'a' {
		var round float64 = float64(8)
		var re int32
		if t&int32(32) != 0 {
			*(*uintptr)(unsafe.Pointer(&prefix)) += uintptr(int32(9))
		}
		pl += int32(2)
		if p < int32(0) || p >= 12 {
			re = int32(0)
		} else {
			re = 12 - p
		}
		if re != 0 {
			round *= float64(2)
			for func() (_cgo_ret int32) {
				_cgo_addr := &re
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				round *= float64(int32(16))
			}
			if int32(*prefix) == '-' {
				y = -y
				y -= round
				y += round
				y = -y
			} else {
				y += round
				y -= round
			}
		}
		estr = _cgos_fmt_u_vfprintf(uint64(func() int32 {
			if e2 < int32(0) {
				return -e2
			} else {
				return e2
			}
		}()), ebuf)
		if uintptr(unsafe.Pointer(estr)) == uintptr(unsafe.Pointer(ebuf)) {
			*func() (_cgo_ret *int8) {
				_cgo_addr := &estr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return *_cgo_addr
			}() = int8('0')
		}
		*func() (_cgo_ret *int8) {
			_cgo_addr := &estr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = int8(func() int32 {
			if e2 < int32(0) {
				return '-'
			} else {
				return '+'
			}
		}())
		*func() (_cgo_ret *int8) {
			_cgo_addr := &estr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = int8(t + 15)
		s = (*int8)(unsafe.Pointer(&buf))
		for {
			var x int32 = int32(y)
			*func() (_cgo_ret *int8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = int8(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&_cgos_xdigits_vfprintf)))) + uintptr(x)))) | t&int32(32))
			y = float64(int32(16)) * (y - float64(x))
			if int64(uintptr(unsafe.Pointer(s))-uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))) == int64(1) && (y != 0 || p > int32(0) || uint32(fl)&8 != 0) {
				*func() (_cgo_ret *int8) {
					_cgo_addr := &s
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() = int8('.')
			}
			if !(y != 0) {
				break
			}
		}
		if int64(p) > int64(2147483645)-int64(uintptr(unsafe.Pointer(ebuf))-uintptr(unsafe.Pointer(estr)))-int64(pl) {
			return -1
		}
		if p != 0 && uintptr(unsafe.Pointer(s))-uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))-uintptr(int64(2)) < uintptr(int64(p)) {
			l = int32(int64(p+int32(2)) + int64(uintptr(unsafe.Pointer(ebuf))-uintptr(unsafe.Pointer(estr))))
		} else {
			l = int32(uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + (uintptr(unsafe.Pointer(ebuf)) - uintptr(unsafe.Pointer(estr))))
		}
		_cgos_pad_vfprintf(f, int8(' '), w, pl+l, fl)
		_cgos_out_vfprintf(f, prefix, uint64(pl))
		_cgos_pad_vfprintf(f, int8('0'), w, pl+l, int32(uint32(fl)^65536))
		_cgos_out_vfprintf(f, (*int8)(unsafe.Pointer(&buf)), uint64(uintptr(unsafe.Pointer(s))-uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))))
		_cgos_pad_vfprintf(f, int8('0'), int32(int64(l)-int64(uintptr(unsafe.Pointer(ebuf))-uintptr(unsafe.Pointer(estr)))-int64(uintptr(unsafe.Pointer(s))-uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))))), int32(0), int32(0))
		_cgos_out_vfprintf(f, estr, uint64(uintptr(unsafe.Pointer(ebuf))-uintptr(unsafe.Pointer(estr))))
		_cgos_pad_vfprintf(f, int8(' '), w, pl+l, int32(uint32(fl)^8192))
		return func() int32 {
			if w > pl+l {
				return w
			} else {
				return pl + l
			}
		}()
	}
	if p < int32(0) {
		p = int32(6)
	}
	if y != 0 {
		func() int32 {
			y *= float64(268435456)
			return func() (_cgo_ret int32) {
				_cgo_addr := &e2
				*_cgo_addr -= int32(28)
				return *_cgo_addr
			}()
		}()
	}
	if e2 < int32(0) {
		a = func() (_cgo_ret *uint32) {
			_cgo_addr := &r
			*_cgo_addr = func() (_cgo_ret *uint32) {
				_cgo_addr := &z
				*_cgo_addr = (*uint32)(unsafe.Pointer(&big))
				return *_cgo_addr
			}()
			return *_cgo_addr
		}()
	} else {
		a = func() (_cgo_ret *uint32) {
			_cgo_addr := &r
			*_cgo_addr = func() (_cgo_ret *uint32) {
				_cgo_addr := &z
				*_cgo_addr = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&big))))+uintptr(504/4)*4))))-uintptr(int32(53))*4)))) - uintptr(int32(1))*4))
				return *_cgo_addr
			}()
			return *_cgo_addr
		}()
	}
	for {
		*z = uint32(y)
		y = float64(int32(1000000000)) * (y - float64(*func() (_cgo_ret *uint32) {
			_cgo_addr := &z
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}()))
		if !(y != 0) {
			break
		}
	}
	for e2 > int32(0) {
		var carry uint32 = uint32(0)
		var sh int32 = func() int32 {
			if int32(29) < e2 {
				return int32(29)
			} else {
				return e2
			}
		}()
		for d = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(z)) - uintptr(int32(1))*4)); uintptr(unsafe.Pointer(d)) >= uintptr(unsafe.Pointer(a)); *(*uintptr)(unsafe.Pointer(&d)) -= 4 {
			var x uint64 = uint64(*d)<<sh + uint64(carry)
			*d = uint32(x % uint64(1000000000))
			carry = uint32(x / uint64(1000000000))
		}
		if carry != 0 {
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &a
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) -= 4
				return *_cgo_addr
			}() = carry
		}
		for uintptr(unsafe.Pointer(z)) > uintptr(unsafe.Pointer(a)) && !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(z)) - uintptr(1)*4)) != 0) {
			*(*uintptr)(unsafe.Pointer(&z)) -= 4
		}
		e2 -= sh
	}
	for e2 < int32(0) {
		var carry uint32 = uint32(0)
		var b *uint32
		var sh int32 = func() int32 {
			if int32(9) < -e2 {
				return int32(9)
			} else {
				return -e2
			}
		}()
		var need int32 = int32(uint32(1) + (uint32(p)+17+uint32(8))/uint32(9))
		for d = a; uintptr(unsafe.Pointer(d)) < uintptr(unsafe.Pointer(z)); *(*uintptr)(unsafe.Pointer(&d)) += 4 {
			var rm uint32 = *d & uint32(int32(1)<<sh-int32(1))
			*d = *d>>sh + carry
			carry = uint32(int32(1000000000)>>sh) * rm
		}
		if !(*a != 0) {
			*(*uintptr)(unsafe.Pointer(&a)) += 4
		}
		if carry != 0 {
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &z
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = carry
		}
		b = func() *uint32 {
			if t|int32(32) == 'f' {
				return r
			} else {
				return a
			}
		}()
		if (uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(b)))*4 > uintptr(int64(need)) {
			z = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + uintptr(need)*4))
		}
		e2 += sh
	}
	if uintptr(unsafe.Pointer(a)) < uintptr(unsafe.Pointer(z)) {
		for func() int32 {
			i = int32(10)
			return func() (_cgo_ret int32) {
				_cgo_addr := &e
				*_cgo_addr = int32(int64(9) * int64((uintptr(unsafe.Pointer(r))-uintptr(unsafe.Pointer(a)))*4))
				return *_cgo_addr
			}()
		}(); *a >= uint32(i); func() int32 {
			i *= int32(10)
			return func() (_cgo_ret int32) {
				_cgo_addr := &e
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}()
		}() {
		}
	} else {
		e = int32(0)
	}
	j = p - func() int32 {
		if t|int32(32) != 'f' {
			return 1
		} else {
			return 0
		}
	}()*e - func() int32 {
		if t|int32(32) == 'g' && p != 0 {
			return 1
		} else {
			return 0
		}
	}()
	if int64(j) < int64(9)*int64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(r)))*4-uintptr(int64(1))) {
		var x uint32
		d = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(r))+uintptr(int32(1))*4)))) + uintptr((j+9216)/int32(9)-int32(1024))*4))
		j += 9216
		j %= int32(9)
		for func() int32 {
			i = int32(10)
			return func() (_cgo_ret int32) {
				_cgo_addr := &j
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}()
		}(); j < int32(9); func() int32 {
			i *= int32(10)
			return func() (_cgo_ret int32) {
				_cgo_addr := &j
				_cgo_ret = *_cgo_addr
				*_cgo_addr++
				return
			}()
		}() {
		}
		x = *d % uint32(i)
		if x != 0 || uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(d))+uintptr(int32(1))*4)))) != uintptr(unsafe.Pointer(z)) {
			var round float64 = float64(int32(2)) / 2.22044604925031308085e-16
			var small float64
			if *d/uint32(i)&uint32(1) != 0 || i == int32(1000000000) && uintptr(unsafe.Pointer(d)) > uintptr(unsafe.Pointer(a)) && *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) - uintptr(1)*4))&uint32(1) != 0 {
				round += float64(int32(2))
			}
			if x < uint32(i/int32(2)) {
				small = float64(0.5)
			} else if x == uint32(i/int32(2)) && uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(d))+uintptr(int32(1))*4)))) == uintptr(unsafe.Pointer(z)) {
				small = float64(1)
			} else {
				small = float64(1.5)
			}
			if pl != 0 && int32(*prefix) == '-' {
				func() float64 {
					round *= float64(-1)
					return func() (_cgo_ret float64) {
						_cgo_addr := &small
						*_cgo_addr *= float64(-1)
						return *_cgo_addr
					}()
				}()
			}
			*d -= x
			if round+small != round {
				*d = *d + uint32(i)
				for *d > uint32(999999999) {
					*func() (_cgo_ret *uint32) {
						_cgo_addr := &d
						_cgo_ret = *_cgo_addr
						*(*uintptr)(unsafe.Pointer(_cgo_addr)) -= 4
						return
					}() = uint32(0)
					if uintptr(unsafe.Pointer(d)) < uintptr(unsafe.Pointer(a)) {
						*func() (_cgo_ret *uint32) {
							_cgo_addr := &a
							*(*uintptr)(unsafe.Pointer(_cgo_addr)) -= 4
							return *_cgo_addr
						}() = uint32(0)
					}
					*d++
				}
				for func() int32 {
					i = int32(10)
					return func() (_cgo_ret int32) {
						_cgo_addr := &e
						*_cgo_addr = int32(int64(9) * int64((uintptr(unsafe.Pointer(r))-uintptr(unsafe.Pointer(a)))*4))
						return *_cgo_addr
					}()
				}(); *a >= uint32(i); func() int32 {
					i *= int32(10)
					return func() (_cgo_ret int32) {
						_cgo_addr := &e
						_cgo_ret = *_cgo_addr
						*_cgo_addr++
						return
					}()
				}() {
				}
			}
		}
		if uintptr(unsafe.Pointer(z)) > uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(d))+uintptr(int32(1))*4)))) {
			z = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(int32(1))*4))
		}
	}
	for ; uintptr(unsafe.Pointer(z)) > uintptr(unsafe.Pointer(a)) && !(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(z)) - uintptr(1)*4)) != 0); *(*uintptr)(unsafe.Pointer(&z)) -= 4 {
	}
	if t|int32(32) == 'g' {
		if !(p != 0) {
			p++
		}
		if p > e && e >= -4 {
			t--
			p -= e + int32(1)
		} else {
			t -= int32(2)
			p--
		}
		if !(uint32(fl)&8 != 0) {
			if uintptr(unsafe.Pointer(z)) > uintptr(unsafe.Pointer(a)) && *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(z)) - uintptr(1)*4)) != 0 {
				for func() int32 {
					i = int32(10)
					return func() (_cgo_ret int32) {
						_cgo_addr := &j
						*_cgo_addr = int32(0)
						return *_cgo_addr
					}()
				}(); *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(z)) - uintptr(1)*4))%uint32(i) == uint32(0); func() int32 {
					i *= int32(10)
					return func() (_cgo_ret int32) {
						_cgo_addr := &j
						_cgo_ret = *_cgo_addr
						*_cgo_addr++
						return
					}()
				}() {
				}
			} else {
				j = int32(9)
			}
			if t|int32(32) == 'f' {
				p = int32(func() int64 {
					if int64(p) < func() int64 {
						if int64(0) > int64(9)*int64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(r)))*4-uintptr(int64(1)))-int64(j) {
							return int64(0)
						} else {
							return int64(9)*int64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(r)))*4-uintptr(int64(1))) - int64(j)
						}
					}() {
						return int64(p)
					} else {
						return func() int64 {
							if int64(0) > int64(9)*int64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(r)))*4-uintptr(int64(1)))-int64(j) {
								return int64(0)
							} else {
								return int64(9)*int64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(r)))*4-uintptr(int64(1))) - int64(j)
							}
						}()
					}
				}())
			} else {
				p = int32(func() int64 {
					if int64(p) < func() int64 {
						if int64(0) > int64(9)*int64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(r)))*4-uintptr(int64(1)))+int64(e)-int64(j) {
							return int64(0)
						} else {
							return int64(9)*int64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(r)))*4-uintptr(int64(1))) + int64(e) - int64(j)
						}
					}() {
						return int64(p)
					} else {
						return func() int64 {
							if int64(0) > int64(9)*int64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(r)))*4-uintptr(int64(1)))+int64(e)-int64(j) {
								return int64(0)
							} else {
								return int64(9)*int64((uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(r)))*4-uintptr(int64(1))) + int64(e) - int64(j)
							}
						}()
					}
				}())
			}
		}
	}
	if p > 2147483646-func() int32 {
		if p != 0 || uint32(fl)&8 != 0 {
			return 1
		} else {
			return 0
		}
	}() {
		return -1
	}
	l = int32(1) + p + func() int32 {
		if p != 0 || uint32(fl)&8 != 0 {
			return 1
		} else {
			return 0
		}
	}()
	if t|int32(32) == 'f' {
		if e > int32(2147483647)-l {
			return -1
		}
		if e > int32(0) {
			l += e
		}
	} else {
		estr = _cgos_fmt_u_vfprintf(uint64(func() int32 {
			if e < int32(0) {
				return -e
			} else {
				return e
			}
		}()), ebuf)
		for uintptr(unsafe.Pointer(ebuf))-uintptr(unsafe.Pointer(estr)) < uintptr(int64(2)) {
			*func() (_cgo_ret *int8) {
				_cgo_addr := &estr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return *_cgo_addr
			}() = int8('0')
		}
		*func() (_cgo_ret *int8) {
			_cgo_addr := &estr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = int8(func() int32 {
			if e < int32(0) {
				return '-'
			} else {
				return '+'
			}
		}())
		*func() (_cgo_ret *int8) {
			_cgo_addr := &estr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = int8(t)
		if uintptr(unsafe.Pointer(ebuf))-uintptr(unsafe.Pointer(estr)) > uintptr(int64(int32(2147483647)-l)) {
			return -1
		}
		l += int32(uintptr(unsafe.Pointer(ebuf)) - uintptr(unsafe.Pointer(estr)))
	}
	if l > int32(2147483647)-pl {
		return -1
	}
	_cgos_pad_vfprintf(f, int8(' '), w, pl+l, fl)
	_cgos_out_vfprintf(f, prefix, uint64(pl))
	_cgos_pad_vfprintf(f, int8('0'), w, pl+l, int32(uint32(fl)^65536))
	if t|int32(32) == 'f' {
		if uintptr(unsafe.Pointer(a)) > uintptr(unsafe.Pointer(r)) {
			a = r
		}
		for d = a; uintptr(unsafe.Pointer(d)) <= uintptr(unsafe.Pointer(r)); *(*uintptr)(unsafe.Pointer(&d)) += 4 {
			var s *int8 = _cgos_fmt_u_vfprintf(uint64(*d), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))+uintptr(int32(9)))))
			if uintptr(unsafe.Pointer(d)) != uintptr(unsafe.Pointer(a)) {
				for uintptr(unsafe.Pointer(s)) > uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) {
					*func() (_cgo_ret *int8) {
						_cgo_addr := &s
						*(*uintptr)(unsafe.Pointer(_cgo_addr))--
						return *_cgo_addr
					}() = int8('0')
				}
			} else if uintptr(unsafe.Pointer(s)) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))+uintptr(int32(9)))))) {
				*func() (_cgo_ret *int8) {
					_cgo_addr := &s
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return *_cgo_addr
				}() = int8('0')
			}
			_cgos_out_vfprintf(f, s, uint64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))+uintptr(int32(9))))))-uintptr(unsafe.Pointer(s))))
		}
		if p != 0 || uint32(fl)&8 != 0 {
			_cgos_out_vfprintf(f, (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})), uint64(1))
		}
		for ; uintptr(unsafe.Pointer(d)) < uintptr(unsafe.Pointer(z)) && p > int32(0); func() int32 {
			*(*uintptr)(unsafe.Pointer(&d)) += 4
			return func() (_cgo_ret int32) {
				_cgo_addr := &p
				*_cgo_addr -= int32(9)
				return *_cgo_addr
			}()
		}() {
			var s *int8 = _cgos_fmt_u_vfprintf(uint64(*d), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))+uintptr(int32(9)))))
			for uintptr(unsafe.Pointer(s)) > uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) {
				*func() (_cgo_ret *int8) {
					_cgo_addr := &s
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return *_cgo_addr
				}() = int8('0')
			}
			_cgos_out_vfprintf(f, s, uint64(func() int32 {
				if int32(9) < p {
					return int32(9)
				} else {
					return p
				}
			}()))
		}
		_cgos_pad_vfprintf(f, int8('0'), p+int32(9), int32(9), int32(0))
	} else {
		if uintptr(unsafe.Pointer(z)) <= uintptr(unsafe.Pointer(a)) {
			z = (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(int32(1))*4))
		}
		for d = a; uintptr(unsafe.Pointer(d)) < uintptr(unsafe.Pointer(z)) && p >= int32(0); *(*uintptr)(unsafe.Pointer(&d)) += 4 {
			var s *int8 = _cgos_fmt_u_vfprintf(uint64(*d), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))+uintptr(int32(9)))))
			if uintptr(unsafe.Pointer(s)) == uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))+uintptr(int32(9)))))) {
				*func() (_cgo_ret *int8) {
					_cgo_addr := &s
					*(*uintptr)(unsafe.Pointer(_cgo_addr))--
					return *_cgo_addr
				}() = int8('0')
			}
			if uintptr(unsafe.Pointer(d)) != uintptr(unsafe.Pointer(a)) {
				for uintptr(unsafe.Pointer(s)) > uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) {
					*func() (_cgo_ret *int8) {
						_cgo_addr := &s
						*(*uintptr)(unsafe.Pointer(_cgo_addr))--
						return *_cgo_addr
					}() = int8('0')
				}
			} else {
				_cgos_out_vfprintf(f, func() (_cgo_ret *int8) {
					_cgo_addr := &s
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}(), uint64(1))
				if p > int32(0) || uint32(fl)&8 != 0 {
					_cgos_out_vfprintf(f, (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})), uint64(1))
				}
			}
			_cgos_out_vfprintf(f, s, uint64(func() int64 {
				if uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))+uintptr(int32(9))))))-uintptr(unsafe.Pointer(s)) < uintptr(int64(p)) {
					return int64(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))+uintptr(int32(9)))))) - uintptr(unsafe.Pointer(s)))
				} else {
					return int64(p)
				}
			}()))
			p -= int32(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf))))+uintptr(int32(9)))))) - uintptr(unsafe.Pointer(s)))
		}
		_cgos_pad_vfprintf(f, int8('0'), p+int32(18), int32(18), int32(0))
		_cgos_out_vfprintf(f, estr, uint64(uintptr(unsafe.Pointer(ebuf))-uintptr(unsafe.Pointer(estr))))
	}
	_cgos_pad_vfprintf(f, int8(' '), w, pl+l, int32(uint32(fl)^8192))
	return func() int32 {
		if w > pl+l {
			return w
		} else {
			return pl + l
		}
	}()
}
func _cgos_getint_vfprintf(s **int8) int32 {
	var i int32
	for i = int32(0); func() int32 {
		if int32(0) != 0 {
			return Isdigit(int32(**s))
		} else {
			return func() int32 {
				if uint32(**s)-uint32('0') < uint32(10) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0; *(*uintptr)(unsafe.Pointer(&*s))++ {
		if uint32(i) > 214748364 || int32(**s)-'0' > int32(2147483647)-int32(10)*i {
			i = -1
		} else {
			i = int32(10)*i + (int32(**s) - '0')
		}
	}
	return i
}
func _cgos_printf_core_vfprintf(f *Struct__IO_FILE, fmt *int8, ap *[]interface {
}, nl_arg *_cgos_arg_vfprintf, nl_type *int32) int32 {
	var (
		_tag_cgo1 int32
		_nm_cgo2  bool
	)
	var a *int8
	var z *int8
	var s *int8 = (*int8)(unsafe.Pointer(fmt))
	var l10n uint32 = uint32(0)
	var fl uint32
	var w int32
	var p int32
	var xp int32
	var arg _cgos_arg_vfprintf
	var argpos int32
	var st uint32
	var ps uint32
	var cnt int32 = int32(0)
	var l int32 = int32(0)
	var i uint64
	var buf [40]int8
	var prefix *int8
	var t int32
	var pl int32
	var wc [2]uint32
	var ws *uint32
	var mb [4]int8
	for {
		if l > int32(2147483647)-cnt {
			goto overflow
		}
		cnt += l
		if !(*s != 0) {
			break
		}
		for a = s; int32(*s) != 0 && int32(*s) != '%'; *(*uintptr)(unsafe.Pointer(&s))++ {
		}
		for z = s; int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(0))))) == '%' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))))) == '%'; func() *int8 {
			*(*uintptr)(unsafe.Pointer(&z))++
			return func() (_cgo_ret *int8) {
				_cgo_addr := &s
				*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(int32(2))
				return *_cgo_addr
			}()
		}() {
		}
		if uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(a)) > uintptr(int64(int32(2147483647)-cnt)) {
			goto overflow
		}
		l = int32(uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(a)))
		if f != nil {
			_cgos_out_vfprintf(f, a, uint64(l))
		}
		if l != 0 {
			continue
		}
		if func() int32 {
			if int32(0) != 0 {
				return Isdigit(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))))))
			} else {
				return func() int32 {
					if uint32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1)))))-uint32('0') < uint32(10) {
						return 1
					} else {
						return 0
					}
				}()
			}
		}() != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2))))) == '$' {
			l10n = uint32(1)
			argpos = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))))) - '0'
			*(*uintptr)(unsafe.Pointer(&s)) += uintptr(int32(3))
		} else {
			argpos = -1
			*(*uintptr)(unsafe.Pointer(&s))++
		}
		for fl = uint32(0); uint32(*s)-uint32(' ') < uint32(32) && 75913&(uint32(1)<<(int32(*s)-' ')) != 0; *(*uintptr)(unsafe.Pointer(&s))++ {
			fl |= uint32(1) << (int32(*s) - ' ')
		}
		if int32(*s) == '*' {
			if func() int32 {
				if int32(0) != 0 {
					return Isdigit(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))))))
				} else {
					return func() int32 {
						if uint32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1)))))-uint32('0') < uint32(10) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}() != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2))))) == '$' {
				l10n = uint32(1)
				*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1)))))-'0')*4)) = INT
				w = int32((*(*_cgos_arg_vfprintf)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_arg)) + uintptr(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1)))))-'0')*8))).i)
				*(*uintptr)(unsafe.Pointer(&s)) += uintptr(int32(3))
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
				*(*uintptr)(unsafe.Pointer(&s))++
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
			*_cgo_addr = _cgos_getint_vfprintf(&s)
			return *_cgo_addr
		}() < int32(0) {
			goto overflow
		}
		if int32(*s) == '.' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(1))))) == '*' {
			if func() int32 {
				if int32(0) != 0 {
					return Isdigit(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2))))))
				} else {
					return func() int32 {
						if uint32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2)))))-uint32('0') < uint32(10) {
							return 1
						} else {
							return 0
						}
					}()
				}
			}() != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(3))))) == '$' {
				*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2)))))-'0')*4)) = INT
				p = int32((*(*_cgos_arg_vfprintf)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_arg)) + uintptr(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(int32(2)))))-'0')*8))).i)
				*(*uintptr)(unsafe.Pointer(&s)) += uintptr(int32(4))
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
				*(*uintptr)(unsafe.Pointer(&s)) += uintptr(int32(2))
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
		} else if int32(*s) == '.' {
			*(*uintptr)(unsafe.Pointer(&s))++
			p = _cgos_getint_vfprintf(&s)
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
			st = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[58]uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*[58]uint8)(unsafe.Pointer(&_cgos_states_vfprintf)))) + uintptr(st)*58)))))) + uintptr(int32(*func() (_cgo_ret *int8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}())-'A'))))
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
			func() _cgos_arg_vfprintf {
				*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(argpos)*4)) = int32(st)
				return func() (_cgo_ret _cgos_arg_vfprintf) {
					_cgo_addr := &arg
					*_cgo_addr = *(*_cgos_arg_vfprintf)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_arg)) + uintptr(argpos)*8))
					return *_cgo_addr
				}()
			}()
		} else if f != nil {
			_cgos_pop_arg_vfprintf(&arg, int32(st), ap)
		} else {
			return int32(0)
		}
		if !(f != nil) {
			continue
		}
		z = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&buf)))) + uintptr(40)))
		prefix = (*int8)(unsafe.Pointer(&[10]int8{'-', '+', ' ', ' ', ' ', '0', 'X', '0', 'x', '\x00'}))
		pl = int32(0)
		t = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) - uintptr(1))))
		if ps != 0 && t&int32(15) == int32(3) {
			t &= -33
		}
		if fl&8192 != 0 {
			fl &= 4294901759
		}
		_tag_cgo1, _nm_cgo2 = t, true
		if _nm_cgo2 && _tag_cgo1 != 'n' {
			goto _cgol_1
		}
		_nm_cgo2 = false
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
	_cgol_1:
		if _nm_cgo2 && _tag_cgo1 != 'p' {
			goto _cgol_2
		}
		_nm_cgo2 = false
		p = int32(func() uint64 {
			if uint64(p) > 16 {
				return uint64(p)
			} else {
				return 16
			}
		}())
		t = int32('x')
		fl |= 8
	_cgol_2:
		if _nm_cgo2 && _tag_cgo1 != 'x' {
			goto _cgol_3
		}
		_nm_cgo2 = false
	_cgol_3:
		if _nm_cgo2 && _tag_cgo1 != 'X' {
			goto _cgol_4
		}
		_nm_cgo2 = false
		a = _cgos_fmt_x_vfprintf(arg.i, z, t&int32(32))
		if arg.i != 0 && fl&8 != 0 {
			func() int32 {
				*(*uintptr)(unsafe.Pointer(&prefix)) += uintptr(t >> int32(4))
				return func() (_cgo_ret int32) {
					_cgo_addr := &pl
					*_cgo_addr = int32(2)
					return *_cgo_addr
				}()
			}()
		}
		if true {
			goto _cgol_5
		}
	_cgol_4:
		if _nm_cgo2 && _tag_cgo1 != 'o' {
			goto _cgol_6
		}
		_nm_cgo2 = false
		a = _cgos_fmt_o_vfprintf(arg.i, z)
		if fl&8 != 0 && int64(p) < int64(uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(a))+uintptr(int64(1))) {
			p = int32(uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(a)) + uintptr(int64(1)))
		}
	_cgol_5:
		if true {
			goto _cgol_7
		}
	_cgol_6:
		if _nm_cgo2 && _tag_cgo1 != 'd' {
			goto _cgol_8
		}
		_nm_cgo2 = false
	_cgol_8:
		if _nm_cgo2 && _tag_cgo1 != 'i' {
			goto _cgol_9
		}
		_nm_cgo2 = false
		pl = int32(1)
		if arg.i > uint64(9223372036854775807) {
			arg.i = -arg.i
		} else if fl&2048 != 0 {
			*(*uintptr)(unsafe.Pointer(&prefix))++
		} else if fl&1 != 0 {
			*(*uintptr)(unsafe.Pointer(&prefix)) += uintptr(int32(2))
		} else {
			pl = int32(0)
		}
	_cgol_9:
		if _nm_cgo2 && _tag_cgo1 != 'u' {
			goto _cgol_10
		}
		_nm_cgo2 = false
		a = _cgos_fmt_u_vfprintf(arg.i, z)
	_cgol_7:
		if xp != 0 && p < int32(0) {
			goto overflow
		}
		if xp != 0 {
			fl &= 4294901759
		}
		if !(arg.i != 0) && !(p != 0) {
			a = z
			goto _cgol_11
		}
		p = int32(func() int64 {
			if int64(p) > int64(uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(a))+uintptr(func() int64 {
				if !(arg.i != 0) {
					return 1
				} else {
					return 0
				}
			}())) {
				return int64(p)
			} else {
				return int64(uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(a)) + uintptr(func() int64 {
					if !(arg.i != 0) {
						return 1
					} else {
						return 0
					}
				}()))
			}
		}())
		goto _cgol_11
	_cgol_10:
		if _nm_cgo2 && _tag_cgo1 != 'c' {
			goto _cgol_12
		}
		_nm_cgo2 = false
		*func() (_cgo_ret *int8) {
			_cgo_addr := &a
			*_cgo_addr = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(z)) - uintptr(func() (_cgo_ret int32) {
				_cgo_addr := &p
				*_cgo_addr = int32(1)
				return *_cgo_addr
			}())))
			return *_cgo_addr
		}() = int8(arg.i)
		fl &= 4294901759
		goto _cgol_11
	_cgol_12:
		if _nm_cgo2 && _tag_cgo1 != 'm' {
			goto _cgol_13
		}
		_nm_cgo2 = false
		if false {
			goto _cgol_15
		}
		a = Strerror(*X__errno_location())
		goto _cgol_14
	_cgol_15:
		;
	_cgol_13:
		if _nm_cgo2 && _tag_cgo1 != 's' {
			goto _cgol_16
		}
		_nm_cgo2 = false
		a = (*int8)(func() unsafe.Pointer {
			if *(*unsafe.Pointer)(unsafe.Pointer(&arg)) != nil {
				return *(*unsafe.Pointer)(unsafe.Pointer(&arg))
			} else {
				return unsafe.Pointer((*int8)(unsafe.Pointer(&[7]int8{'(', 'n', 'u', 'l', 'l', ')', '\x00'})))
			}
		}())
	_cgol_14:
		z = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(Strnlen(a, uint64(func() int32 {
			if p < int32(0) {
				return int32(2147483647)
			} else {
				return p
			}
		}())))))
		if p < int32(0) && int32(*z) != 0 {
			goto overflow
		}
		p = int32(uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(a)))
		fl &= 4294901759
		goto _cgol_11
	_cgol_16:
		if _nm_cgo2 && _tag_cgo1 != 'C' {
			goto _cgol_17
		}
		_nm_cgo2 = false
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&wc)))) + uintptr(int32(0))*4)) = uint32(arg.i)
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&wc)))) + uintptr(int32(1))*4)) = uint32(0)
		*(*unsafe.Pointer)(unsafe.Pointer(&arg)) = unsafe.Pointer((*uint32)(unsafe.Pointer(&wc)))
		p = -1
	_cgol_17:
		if _nm_cgo2 && _tag_cgo1 != 'S' {
			goto _cgol_18
		}
		_nm_cgo2 = false
		ws = (*uint32)(*(*unsafe.Pointer)(unsafe.Pointer(&arg)))
		for i = uint64(func() (_cgo_ret int32) {
			_cgo_addr := &l
			*_cgo_addr = int32(0)
			return *_cgo_addr
		}()); i < uint64(p) && *ws != 0 && func() (_cgo_ret int32) {
			_cgo_addr := &l
			*_cgo_addr = Wctomb((*int8)(unsafe.Pointer(&mb)), *func() (_cgo_ret *uint32) {
				_cgo_addr := &ws
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}())
			return *_cgo_addr
		}() >= int32(0) && uint64(l) <= uint64(p)-i; i += uint64(l) {
		}
		if l < int32(0) {
			return -1
		}
		if i > uint64(2147483647) {
			goto overflow
		}
		p = int32(i)
		_cgos_pad_vfprintf(f, int8(' '), w, p, int32(fl))
		ws = (*uint32)(*(*unsafe.Pointer)(unsafe.Pointer(&arg)))
		for i = uint64(0); i < uint64(uint32(0)+uint32(p)) && *ws != 0 && i+uint64(func() (_cgo_ret int32) {
			_cgo_addr := &l
			*_cgo_addr = Wctomb((*int8)(unsafe.Pointer(&mb)), *func() (_cgo_ret *uint32) {
				_cgo_addr := &ws
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}())
			return *_cgo_addr
		}()) <= uint64(p); i += uint64(l) {
			_cgos_out_vfprintf(f, (*int8)(unsafe.Pointer(&mb)), uint64(l))
		}
		_cgos_pad_vfprintf(f, int8(' '), w, p, int32(fl^8192))
		l = func() int32 {
			if w > p {
				return w
			} else {
				return p
			}
		}()
		continue
	_cgol_18:
		if _nm_cgo2 && _tag_cgo1 != 'e' {
			goto _cgol_19
		}
		_nm_cgo2 = false
	_cgol_19:
		if _nm_cgo2 && _tag_cgo1 != 'f' {
			goto _cgol_20
		}
		_nm_cgo2 = false
	_cgol_20:
		if _nm_cgo2 && _tag_cgo1 != 'g' {
			goto _cgol_21
		}
		_nm_cgo2 = false
	_cgol_21:
		if _nm_cgo2 && _tag_cgo1 != 'a' {
			goto _cgol_22
		}
		_nm_cgo2 = false
	_cgol_22:
		if _nm_cgo2 && _tag_cgo1 != 'E' {
			goto _cgol_23
		}
		_nm_cgo2 = false
	_cgol_23:
		if _nm_cgo2 && _tag_cgo1 != 'F' {
			goto _cgol_24
		}
		_nm_cgo2 = false
	_cgol_24:
		if _nm_cgo2 && _tag_cgo1 != 'G' {
			goto _cgol_25
		}
		_nm_cgo2 = false
	_cgol_25:
		if _nm_cgo2 && _tag_cgo1 != 'A' {
			goto _cgol_26
		}
		_nm_cgo2 = false
		if xp != 0 && p < int32(0) {
			goto overflow
		}
		l = _cgos_fmt_fp_vfprintf(f, *(*float64)(unsafe.Pointer(&arg)), w, p, int32(fl), t)
		if l < int32(0) {
			goto overflow
		}
		continue
		goto _cgol_11
	_cgol_26:
		;
	_cgol_11:
		if int64(p) < int64(uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(a))) {
			p = int32(uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(a)))
		}
		if p > int32(2147483647)-pl {
			goto overflow
		}
		if w < pl+p {
			w = pl + p
		}
		if w > int32(2147483647)-cnt {
			goto overflow
		}
		_cgos_pad_vfprintf(f, int8(' '), w, pl+p, int32(fl))
		_cgos_out_vfprintf(f, prefix, uint64(pl))
		_cgos_pad_vfprintf(f, int8('0'), w, pl+p, int32(fl^65536))
		_cgos_pad_vfprintf(f, int8('0'), p, int32(uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(a))), int32(0))
		_cgos_out_vfprintf(f, a, uint64(uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(a))))
		_cgos_pad_vfprintf(f, int8(' '), w, pl+p, int32(fl^8192))
		l = w
	}
	if f != nil {
		return cnt
	}
	if !(l10n != 0) {
		return int32(0)
	}
	for i = uint64(1); i <= uint64(9) && *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(i)*4)) != 0; i++ {
		_cgos_pop_arg_vfprintf((*_cgos_arg_vfprintf)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_arg))+uintptr(i)*8)), *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(i)*4)), ap)
	}
	for ; i <= uint64(9) && !(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(nl_type)) + uintptr(i)*4)) != 0); i++ {
	}
	if i <= uint64(9) {
		goto inval
	}
	return int32(1)
inval:
	*X__errno_location() = int32(22)
	return -1
overflow:
	*X__errno_location() = int32(75)
	return -1
}
func Vfprintf(f *Struct__IO_FILE, fmt *int8, ap []interface {
}) int32 {
	var ap2 []interface {
	}
	var nl_type [10]int32 = [10]int32{int32(0)}
	var nl_arg [10]_cgos_arg_vfprintf
	var internal_buf [80]uint8
	var saved_buf *uint8 = nil
	var olderr int32
	var ret int32
	ap2 = ap
	if _cgos_printf_core_vfprintf(nil, fmt, &ap2, (*_cgos_arg_vfprintf)(unsafe.Pointer(&nl_arg)), (*int32)(unsafe.Pointer(&nl_type))) < int32(0) {
		return -1
	}
	var __need_unlock int32 = func() int32 {
		if f.Lock >= int32(0) {
			return __lockfile(f)
		} else {
			return int32(0)
		}
	}()
	olderr = int32(f.Flags & uint32(32))
	if f.Mode < int32(1) {
		f.Flags &= uint32(4294967263)
	}
	if !(f.Buf_size != 0) {
		saved_buf = f.Buf
		f.Buf = (*uint8)(unsafe.Pointer(&internal_buf))
		f.Buf_size = uint64(80)
		f.Wpos = func() (_cgo_ret *uint8) {
			_cgo_addr := &f.Wbase
			*_cgo_addr = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Wend
				*_cgo_addr = (*uint8)(nil)
				return *_cgo_addr
			}()
			return *_cgo_addr
		}()
	}
	if !(f.Wend != nil) && __towrite(f) != 0 {
		ret = -1
	} else {
		ret = _cgos_printf_core_vfprintf(f, fmt, &ap2, (*_cgos_arg_vfprintf)(unsafe.Pointer(&nl_arg)), (*int32)(unsafe.Pointer(&nl_type)))
	}
	if saved_buf != nil {
		f.Write(f, nil, uint64(0))
		if !(f.Wpos != nil) {
			ret = -1
		}
		f.Buf = saved_buf
		f.Buf_size = uint64(0)
		f.Wpos = func() (_cgo_ret *uint8) {
			_cgo_addr := &f.Wbase
			*_cgo_addr = func() (_cgo_ret *uint8) {
				_cgo_addr := &f.Wend
				*_cgo_addr = (*uint8)(nil)
				return *_cgo_addr
			}()
			return *_cgo_addr
		}()
	}
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
