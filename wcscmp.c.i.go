package libc

import unsafe "unsafe"

func Wcscmp(l *int32, r *int32) int32 {
	for ; *l == *r && *l != 0 && *r != 0; func() *int32 {
		*(*uintptr)(unsafe.Pointer(&l)) += 4
		return func() (_cgo_ret *int32) {
			_cgo_addr := &r
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}()
	}() {
	}
	return *l - *r
}
