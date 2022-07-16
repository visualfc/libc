package libc

import unsafe "unsafe"

func Strtok(s *int8, sep *int8) *int8 {
	if !(s != nil) && !(func() (_cgo_ret *int8) {
		_cgo_addr := &s
		*_cgo_addr = p_cgo689
		return *_cgo_addr
	}() != nil) {
		return (*int8)(nil)
	}
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(Strspn(s, sep))
	if !(*s != 0) {
		return func() (_cgo_ret *int8) {
			_cgo_addr := &p_cgo689
			*_cgo_addr = (*int8)(nil)
			return *_cgo_addr
		}()
	}
	p_cgo689 = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(Strcspn(s, sep))))
	if *p_cgo689 != 0 {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &p_cgo689
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = int8(0)
	} else {
		p_cgo689 = (*int8)(nil)
	}
	return s
}

var p_cgo689 *int8
