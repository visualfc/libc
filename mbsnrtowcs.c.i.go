package libc

import unsafe "unsafe"

func mbsnrtowcs(wcs *int32, src **int8, n uint64, wn uint64, st *Struct___mbstate_t) uint64 {
	var l uint64
	var cnt uint64 = uint64(0)
	var n2 uint64
	var ws *int32
	var wbuf [256]int32
	var s *int8 = *src
	var tmp_s *int8
	if !(wcs != nil) {
		func() uint64 {
			ws = (*int32)(unsafe.Pointer(&wbuf))
			return func() (_cgo_ret uint64) {
				_cgo_addr := &wn
				*_cgo_addr = uint64(256)
				return *_cgo_addr
			}()
		}()
	} else {
		ws = wcs
	}
	for s != nil && wn != 0 && (func() (_cgo_ret uint64) {
		_cgo_addr := &n2
		*_cgo_addr = n / uint64(4)
		return *_cgo_addr
	}() >= wn || n2 > uint64(32)) {
		if n2 >= wn {
			n2 = wn
		}
		tmp_s = s
		l = mbsrtowcs(ws, &s, n2, st)
		if !(l+uint64(1) != 0) {
			cnt = l
			wn = uint64(0)
			break
		}
		if uintptr(unsafe.Pointer(ws)) != uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&wbuf)))) {
			*(*uintptr)(unsafe.Pointer(&ws)) += uintptr(l) * 4
			wn -= l
		}
		n = func() uint64 {
			if s != nil {
				return n - uint64(uintptr(unsafe.Pointer(s))-uintptr(unsafe.Pointer(tmp_s)))
			} else {
				return uint64(0)
			}
		}()
		cnt += l
	}
	if s != nil {
		for wn != 0 && n != 0 {
			l = mbrtowc(ws, s, n, st)
			if l+uint64(2) <= uint64(2) {
				if !(l+uint64(1) != 0) {
					cnt = l
					break
				}
				if !(l != 0) {
					s = (*int8)(nil)
					break
				}
				*(*uint32)(unsafe.Pointer(st)) = uint32(0)
				break
			}
			*(*uintptr)(unsafe.Pointer(&s)) += uintptr(l)
			n -= l
			*(*uintptr)(unsafe.Pointer(&ws)) += 4
			wn--
			cnt++
		}
	}
	if wcs != nil {
		*src = s
	}
	return cnt
}
