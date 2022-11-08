package libc

import unsafe "unsafe"

func wmemchr(s *int32, c int32, n uint64) *int32 {
	for ; n != 0 && *s != c; func() *int32 {
		n--
		return func() (_cgo_ret *int32) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}()
	}() {
	}
	return func() *int32 {
		if n != 0 {
			return (*int32)(unsafe.Pointer(s))
		} else {
			return nil
		}
	}()
}
