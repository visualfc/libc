package libc

import unsafe "unsafe"

var _cgos_table_intscan [257]uint8 = [257]uint8{uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(0), uint8(1), uint8(2), uint8(3), uint8(4), uint8(5), uint8(6), uint8(7), uint8(8), uint8(9), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(10), uint8(11), uint8(12), uint8(13), uint8(14), uint8(15), uint8(16), uint8(17), uint8(18), uint8(19), uint8(20), uint8(21), uint8(22), uint8(23), uint8(24), uint8(25), uint8(26), uint8(27), uint8(28), uint8(29), uint8(30), uint8(31), uint8(32), uint8(33), uint8(34), uint8(35), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(10), uint8(11), uint8(12), uint8(13), uint8(14), uint8(15), uint8(16), uint8(17), uint8(18), uint8(19), uint8(20), uint8(21), uint8(22), uint8(23), uint8(24), uint8(25), uint8(26), uint8(27), uint8(28), uint8(29), uint8(30), uint8(31), uint8(32), uint8(33), uint8(34), uint8(35), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255), uint8(255)}

func __intscan(f *Struct__IO_FILE, base uint32, pok int32, lim uint64) uint64 {
	var val *uint8 = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_table_intscan)))) + uintptr(int32(1))))
	var c int32
	var neg int32 = int32(0)
	var x uint32
	var y uint64
	if base > uint32(36) || base == uint32(1) {
		*X__errno_location() = int32(22)
		return uint64(0)
	}
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
	if c == '+' || c == '-' {
		neg = -func() int32 {
			if c == '-' {
				return 1
			} else {
				return 0
			}
		}()
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
	if (base == uint32(0) || base == uint32(16)) && c == '0' {
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
		if c|int32(32) == 'x' {
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
			if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c)))) >= int32(16) {
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
				if pok != 0 {
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
				} else {
					__shlim(f, int64(0))
				}
				return uint64(0)
			}
			base = uint32(16)
		} else if base == uint32(0) {
			base = uint32(8)
		}
	} else {
		if base == uint32(0) {
			base = uint32(10)
		}
		if uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c)))) >= base {
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
			__shlim(f, int64(0))
			*X__errno_location() = int32(22)
			return uint64(0)
		}
	}
	if base == uint32(10) {
		for x = uint32(0); uint32(c-'0') < uint32(10) && x <= 429496728; c = func() int32 {
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
		}() {
			x = x*uint32(10) + uint32(c-'0')
		}
		for y = uint64(x); uint32(c-'0') < uint32(10) && y <= 1844674407370955161 && uint64(10)*y <= 18446744073709551615-uint64(c-'0'); c = func() int32 {
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
		}() {
			y = y*uint64(10) + uint64(c-'0')
		}
		if uint32(c-'0') >= uint32(10) {
			goto done
		}
	} else if !(base&(base-uint32(1)) != 0) {
		var bs int32 = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'\x00', '\x01', '\x02', '\x04', '\a', '\x03', '\x06', '\x05', '\x00'})))) + uintptr(uint32(23)*base>>int32(5)&uint32(7)))))
		for x = uint32(0); uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c)))) < base && x <= 134217727; c = func() int32 {
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
		}() {
			x = x<<bs | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c))))
		}
		for y = uint64(x); uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c)))) < base && y <= 18446744073709551615>>bs; c = func() int32 {
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
		}() {
			y = y<<bs | uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c))))
		}
	} else {
		for x = uint32(0); uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c)))) < base && x <= 119304646; c = func() int32 {
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
		}() {
			x = x*base + uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c))))
		}
		for y = uint64(x); uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c)))) < base && y <= 18446744073709551615/uint64(base) && uint64(base)*y <= 18446744073709551615-uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c)))); c = func() int32 {
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
		}() {
			y = y*uint64(base) + uint64(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c))))
		}
	}
	if uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c)))) < base {
		for ; uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(val)) + uintptr(c)))) < base; c = func() int32 {
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
		}() {
		}
		*X__errno_location() = int32(34)
		y = lim
		if lim&uint64(1) != 0 {
			neg = int32(0)
		}
	}
done:
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
	if y >= lim {
		if !(lim&uint64(1) != 0) && !(neg != 0) {
			*X__errno_location() = int32(34)
			return lim - uint64(1)
		} else if y > lim {
			*X__errno_location() = int32(34)
			return lim
		}
	}
	return y ^ uint64(neg) - uint64(neg)
}
