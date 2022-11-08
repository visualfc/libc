package libc

import unsafe "unsafe"

func Wcsncpy(d *int32, s *int32, n uint64) *int32 {
	var a *int32 = d
	for n != 0 && *s != 0 {
		func() int32 {
			n--
			return func() (_cgo_ret int32) {
				_cgo_addr := &*func() (_cgo_ret *int32) {
					_cgo_addr := &d
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
					return
				}()
				*_cgo_addr = *func() (_cgo_ret *int32) {
					_cgo_addr := &s
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
					return
				}()
				return *_cgo_addr
			}()
		}()
	}
	wmemset(d, int32(0), n)
	return a
}
