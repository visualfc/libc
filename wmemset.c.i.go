package libc

import unsafe "unsafe"

func wmemset(d *int32, c int32, n uint64) *int32 {
	var ret *int32 = d
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
		}() = c
	}
	return ret
}
