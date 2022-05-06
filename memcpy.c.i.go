package libc

import unsafe "unsafe"

func memcpy(dest unsafe.Pointer, src unsafe.Pointer, n uint64) unsafe.Pointer {
	var d *uint8 = (*uint8)(dest)
	var s *uint8 = (*uint8)(src)
	type u32 = uint32
	var w uint32
	var x uint32
	for ; uint64(uintptr(unsafe.Pointer(s)))%uint64(4) != 0 && n != 0; n-- {
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}
	if uint64(uintptr(unsafe.Pointer(d)))%uint64(4) == uint64(0) {
		for ; n >= uint64(16); func() uint64 {
			func() *uint8 {
				*(*uintptr)(unsafe.Pointer(&s)) += uintptr(16)
				return func() (_cgo_ret *uint8) {
					_cgo_addr := &d
					*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(16)
					return *_cgo_addr
				}()
			}()
			return func() (_cgo_ret uint64) {
				_cgo_addr := &n
				*_cgo_addr -= uint64(16)
				return *_cgo_addr
			}()
		}() {
			*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(0))))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(0)))))
			*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(4))))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(4)))))
			*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(8))))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(8)))))
			*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(12))))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(12)))))
		}
		if n&uint64(8) != 0 {
			*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(0))))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(0)))))
			*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(4))))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(4)))))
			*(*uintptr)(unsafe.Pointer(&d)) += uintptr(8)
			*(*uintptr)(unsafe.Pointer(&s)) += uintptr(8)
		}
		if n&uint64(4) != 0 {
			*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(0))))) = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(0)))))
			*(*uintptr)(unsafe.Pointer(&d)) += uintptr(4)
			*(*uintptr)(unsafe.Pointer(&s)) += uintptr(4)
		}
		if n&uint64(2) != 0 {
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}
		if n&uint64(1) != 0 {
			*d = *s
		}
		return dest
	}
	if n >= uint64(32) {
		switch uint64(uintptr(unsafe.Pointer(d))) % uint64(4) {
		case uint64(1):
			w = *(*uint32)(unsafe.Pointer(s))
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
			n -= uint64(3)
			for ; n >= uint64(17); func() uint64 {
				func() *uint8 {
					*(*uintptr)(unsafe.Pointer(&s)) += uintptr(16)
					return func() (_cgo_ret *uint8) {
						_cgo_addr := &d
						*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(16)
						return *_cgo_addr
					}()
				}()
				return func() (_cgo_ret uint64) {
					_cgo_addr := &n
					*_cgo_addr -= uint64(16)
					return *_cgo_addr
				}()
			}() {
				x = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(1)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(0))))) = w>>24 | x<<8
				w = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(5)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(4))))) = x>>24 | w<<8
				x = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(9)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(8))))) = w>>24 | x<<8
				w = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(13)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(12))))) = x>>24 | w<<8
			}
			break
		case uint64(2):
			w = *(*uint32)(unsafe.Pointer(s))
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
			n -= uint64(2)
			for ; n >= uint64(18); func() uint64 {
				func() *uint8 {
					*(*uintptr)(unsafe.Pointer(&s)) += uintptr(16)
					return func() (_cgo_ret *uint8) {
						_cgo_addr := &d
						*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(16)
						return *_cgo_addr
					}()
				}()
				return func() (_cgo_ret uint64) {
					_cgo_addr := &n
					*_cgo_addr -= uint64(16)
					return *_cgo_addr
				}()
			}() {
				x = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(2)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(0))))) = w>>16 | x<<16
				w = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(6)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(4))))) = x>>16 | w<<16
				x = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(10)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(8))))) = w>>16 | x<<16
				w = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(14)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(12))))) = x>>16 | w<<16
			}
			break
		case uint64(3):
			w = *(*uint32)(unsafe.Pointer(s))
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &d
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &s
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
			n -= uint64(1)
			for ; n >= uint64(19); func() uint64 {
				func() *uint8 {
					*(*uintptr)(unsafe.Pointer(&s)) += uintptr(16)
					return func() (_cgo_ret *uint8) {
						_cgo_addr := &d
						*(*uintptr)(unsafe.Pointer(&*_cgo_addr)) += uintptr(16)
						return *_cgo_addr
					}()
				}()
				return func() (_cgo_ret uint64) {
					_cgo_addr := &n
					*_cgo_addr -= uint64(16)
					return *_cgo_addr
				}()
			}() {
				x = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(3)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(0))))) = w>>8 | x<<24
				w = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(7)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(4))))) = x>>8 | w<<24
				x = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(11)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(8))))) = w>>8 | x<<24
				w = *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(15)))))
				*(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(d)) + uintptr(12))))) = x>>8 | w<<24
			}
			break
		}
	}
	if n&uint64(16) != 0 {
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}
	if n&uint64(8) != 0 {
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}
	if n&uint64(4) != 0 {
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}
	if n&uint64(2) != 0 {
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}
	if n&uint64(1) != 0 {
		*d = *s
	}
	return dest
	for ; n != 0; n-- {
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *uint8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}
	return dest
}
