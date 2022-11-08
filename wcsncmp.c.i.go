package libc

import unsafe "unsafe"

func Wcsncmp(l *int32, r *int32, n uint64) int32 {
	for ; n != 0 && *l == *r && *l != 0 && *r != 0; func() *int32 {
		func() *int32 {
			n--
			return func() (_cgo_ret *int32) {
				_cgo_addr := &l
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
		}()
		return func() (_cgo_ret *int32) {
			_cgo_addr := &r
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}()
	}() {
	}
	return func() int32 {
		if n != 0 {
			return *l - *r
		} else {
			return int32(0)
		}
	}()
}
