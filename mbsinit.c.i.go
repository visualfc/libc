package libc

import unsafe "unsafe"

func mbsinit(st *Struct___mbstate_t) int32 {
	return func() int32 {
		if !(st != nil) || !(*(*uint32)(unsafe.Pointer(st)) != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
