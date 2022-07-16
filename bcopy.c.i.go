package libc

import unsafe "unsafe"

func bcopy(s1 unsafe.Pointer, s2 unsafe.Pointer, n uint64) {
	memmove(s2, s1, n)
}
