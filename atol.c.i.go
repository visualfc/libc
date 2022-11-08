package libc

import unsafe "unsafe"

func Atol(s *int8) int64 {
	var n int64 = int64(0)
	var neg int32 = int32(0)
	for X__isspace(int32(*s)) != 0 {
		*(*uintptr)(unsafe.Pointer(&s))++
	}
	switch int32(*s) {
	case '-':
		neg = int32(1)
	case '+':
		*(*uintptr)(unsafe.Pointer(&s))++
	}
	for func() int32 {
		if int32(0) != 0 {
			return Isdigit(int32(*s))
		} else {
			return func() int32 {
				if uint32(*s)-uint32('0') < uint32(10) {
					return 1
				} else {
					return 0
				}
			}()
		}
	}() != 0 {
		n = int64(10)*n - int64(int32(*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}())-'0')
	}
	return func() int64 {
		if neg != 0 {
			return n
		} else {
			return -n
		}
	}()
}
