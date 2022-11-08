package libc

import unsafe "unsafe"

func iswspace(wc uint32) int32 {
	return func() int32 {
		if wc != 0 && Wcschr((*int32)(unsafe.Pointer(&_cgos_iswspace_spaces_iswspace)), int32(wc)) != nil {
			return 1
		} else {
			return 0
		}
	}()
}

var _cgos_iswspace_spaces_iswspace [22]int32 = [22]int32{' ', '\t', '\n', '\r', int32(11), int32(12), int32(133), int32(8192), int32(8193), int32(8194), int32(8195), int32(8196), int32(8197), int32(8198), int32(8200), int32(8201), int32(8202), int32(8232), int32(8233), int32(8287), int32(12288), int32(0)}

func __iswspace_l(c uint32, l *Struct___locale_struct) int32 {
	return iswspace(c)
}
