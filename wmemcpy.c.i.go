package libc

import unsafe "unsafe"

func wmemcpy(d *int32, s *int32, n uint64) *int32 {
	var a *int32 = d
	for func() (_cgo_ret uint64) {
		_cgo_addr := &n
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *int32) {
			_cgo_addr := &d
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() = *func() (_cgo_ret *int32) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}()
	}
	return a
}
