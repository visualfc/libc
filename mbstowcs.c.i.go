package libc

import unsafe "unsafe"

func Mbstowcs(ws *int32, s *int8, wn uint64) uint64 {
	return mbsrtowcs(ws, (**int8)(unsafe.Pointer(&s)), wn, nil)
}
