package libc

import unsafe "unsafe"

func Wcsncasecmp(l *int32, r *int32, n uint64) int32 {
	if !(func() (_cgo_ret uint64) {
		_cgo_addr := &n
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0) {
		return int32(0)
	}
	for ; *l != 0 && *r != 0 && n != 0 && (*l == *r || towlower(uint32(*l)) == towlower(uint32(*r))); func() uint64 {
		func() *int32 {
			*(*uintptr)(unsafe.Pointer(&l)) += 4
			return func() (_cgo_ret *int32) {
				_cgo_addr := &r
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
		}()
		return func() (_cgo_ret uint64) {
			_cgo_addr := &n
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}()
	}() {
	}
	return int32(towlower(uint32(*l)) - towlower(uint32(*r)))
}
