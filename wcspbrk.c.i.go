package libc

import unsafe "unsafe"

func Wcspbrk(s *int32, b *int32) *int32 {
	*(*uintptr)(unsafe.Pointer(&s)) += uintptr(Wcscspn(s, b)) * 4
	return func() *int32 {
		if *s != 0 {
			return (*int32)(unsafe.Pointer(s))
		} else {
			return nil
		}
	}()
}
