package libc

import unsafe "unsafe"

func Wcscspn(s *int32, c *int32) uint64 {
	var a *int32
	if !(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(int32(0))*4)) != 0) {
		return Wcslen(s)
	}
	if !(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + uintptr(int32(1))*4)) != 0) {
		return func() uint64 {
			if func() (_cgo_ret *int32) {
				_cgo_addr := &s
				*_cgo_addr = Wcschr(func() (_cgo_ret *int32) {
					_cgo_addr := &a
					*_cgo_addr = s
					return *_cgo_addr
				}(), *c)
				return *_cgo_addr
			}() != nil {
				return uint64((uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer(a))) / 4)
			} else {
				return Wcslen(a)
			}
		}()
	}
	for a = s; *s != 0 && !(Wcschr(c, *s) != nil); *(*uintptr)(unsafe.Pointer(&s)) += 4 {
	}
	return uint64((uintptr(unsafe.Pointer(s)) - uintptr(unsafe.Pointer(a))) / 4)
}
