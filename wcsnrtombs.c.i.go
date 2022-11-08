package libc

import unsafe "unsafe"

func wcsnrtombs(dst *int8, wcs **int32, wn uint64, n uint64, st *Struct___mbstate_t) uint64 {
	var ws *int32 = *wcs
	var cnt uint64 = uint64(0)
	if !(dst != nil) {
		n = uint64(0)
	}
	for ws != nil && wn != 0 {
		var tmp [4]int8
		var l uint64 = wcrtomb(func() *int8 {
			if n < uint64(4) {
				return (*int8)(unsafe.Pointer(&tmp))
			} else {
				return dst
			}
		}(), *ws, nil)
		if l == uint64(18446744073709551615) {
			cnt = uint64(18446744073709551615)
			break
		}
		if dst != nil {
			if n < uint64(4) {
				if l > n {
					break
				}
				Memcpy(unsafe.Pointer(dst), unsafe.Pointer((*int8)(unsafe.Pointer(&tmp))), l)
			}
			*(*uintptr)(unsafe.Pointer(&dst)) += uintptr(l)
			n -= l
		}
		if !(*ws != 0) {
			ws = (*int32)(nil)
			break
		}
		*(*uintptr)(unsafe.Pointer(&ws)) += 4
		wn--
		cnt += l
	}
	if dst != nil {
		*wcs = ws
	}
	return cnt
}
