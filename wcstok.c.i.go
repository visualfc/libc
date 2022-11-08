package libc

import unsafe "unsafe"

func Wcstok(s *int32, sep *int32, p **int32) *int32 {
	if !(s != nil) && !(func() (_cgo_ret *int32) {
		_cgo_addr := &s
		*_cgo_addr = *p
		return *_cgo_addr
	}() != nil) {
		return (*int32)(nil)
	}
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(Wcsspn(s, sep)) * 4
	if !(*s != 0) {
		return func() (_cgo_ret *int32) {
			_cgo_addr := &*p
			*_cgo_addr = (*int32)(nil)
			return *_cgo_addr
		}()
	}
	*p = (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(Wcscspn(s, sep))*4))
	if **p != 0 {
		*func() (_cgo_ret *int32) {
			_cgo_addr := &*p
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() = int32(0)
	} else {
		*p = (*int32)(nil)
	}
	return s
}
