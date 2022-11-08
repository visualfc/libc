package libc

import unsafe "unsafe"

var _cgos_c_time_langinfo [316]int8 = [316]int8{'S', 'u', 'n', '\x00', 'M', 'o', 'n', '\x00', 'T', 'u', 'e', '\x00', 'W', 'e', 'd', '\x00', 'T', 'h', 'u', '\x00', 'F', 'r', 'i', '\x00', 'S', 'a', 't', '\x00', 'S', 'u', 'n', 'd', 'a', 'y', '\x00', 'M', 'o', 'n', 'd', 'a', 'y', '\x00', 'T', 'u', 'e', 's', 'd', 'a', 'y', '\x00', 'W', 'e', 'd', 'n', 'e', 's', 'd', 'a', 'y', '\x00', 'T', 'h', 'u', 'r', 's', 'd', 'a', 'y', '\x00', 'F', 'r', 'i', 'd', 'a', 'y', '\x00', 'S', 'a', 't', 'u', 'r', 'd', 'a', 'y', '\x00', 'J', 'a', 'n', '\x00', 'F', 'e', 'b', '\x00', 'M', 'a', 'r', '\x00', 'A', 'p', 'r', '\x00', 'M', 'a', 'y', '\x00', 'J', 'u', 'n', '\x00', 'J', 'u', 'l', '\x00', 'A', 'u', 'g', '\x00', 'S', 'e', 'p', '\x00', 'O', 'c', 't', '\x00', 'N', 'o', 'v', '\x00', 'D', 'e', 'c', '\x00', 'J', 'a', 'n', 'u', 'a', 'r', 'y', '\x00', 'F', 'e', 'b', 'r', 'u', 'a', 'r', 'y', '\x00', 'M', 'a', 'r', 'c', 'h', '\x00', 'A', 'p', 'r', 'i', 'l', '\x00', 'M', 'a', 'y', '\x00', 'J', 'u', 'n', 'e', '\x00', 'J', 'u', 'l', 'y', '\x00', 'A', 'u', 'g', 'u', 's', 't', '\x00', 'S', 'e', 'p', 't', 'e', 'm', 'b', 'e', 'r', '\x00', 'O', 'c', 't', 'o', 'b', 'e', 'r', '\x00', 'N', 'o', 'v', 'e', 'm', 'b', 'e', 'r', '\x00', 'D', 'e', 'c', 'e', 'm', 'b', 'e', 'r', '\x00', 'A', 'M', '\x00', 'P', 'M', '\x00', '%', 'a', ' ', '%', 'b', ' ', '%', 'e', ' ', '%', 'T', ' ', '%', 'Y', '\x00', '%', 'm', '/', '%', 'd', '/', '%', 'y', '\x00', '%', 'H', ':', '%', 'M', ':', '%', 'S', '\x00', '%', 'I', ':', '%', 'M', ':', '%', 'S', ' ', '%', 'p', '\x00', '\x00', '\x00', '%', 'm', '/', '%', 'd', '/', '%', 'y', '\x00', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '\x00', '%', 'a', ' ', '%', 'b', ' ', '%', 'e', ' ', '%', 'T', ' ', '%', 'Y', '\x00', '%', 'H', ':', '%', 'M', ':', '%', 'S', '\x00'}
var _cgos_c_messages_langinfo [19]int8 = [19]int8{'^', '[', 'y', 'Y', ']', '\x00', '^', '[', 'n', 'N', ']', '\x00', 'y', 'e', 's', '\x00', 'n', 'o', '\x00'}
var _cgos_c_numeric_langinfo [3]int8 = [3]int8{'.', '\x00', '\x00'}

func __nl_langinfo_l(item int32, loc *Struct___locale_struct) *int8 {
	var cat int32 = item >> int32(16)
	var idx int32 = item & int32(65535)
	var str *int8
	if item == int32(14) {
		return func() *int8 {
			if *(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&loc.Cat)))) + uintptr(int32(0))*8)) != nil {
				return (*int8)(unsafe.Pointer(&[6]int8{'U', 'T', 'F', '-', '8', '\x00'}))
			} else {
				return (*int8)(unsafe.Pointer(&[6]int8{'A', 'S', 'C', 'I', 'I', '\x00'}))
			}
		}()
	}
	if idx == int32(65535) && cat < int32(6) {
		return func() *int8 {
			if *(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&loc.Cat)))) + uintptr(cat)*8)) != nil {
				return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&(*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&loc.Cat)))) + uintptr(cat)*8))).name))))
			} else {
				return (*int8)(unsafe.Pointer(&[2]int8{'C', '\x00'}))
			}
		}()
	}
	switch cat {
	case int32(1):
		if idx > int32(1) {
			return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		}
		str = (*int8)(unsafe.Pointer(&_cgos_c_numeric_langinfo))
		break
	case int32(2):
		if idx > int32(49) {
			return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		}
		str = (*int8)(unsafe.Pointer(&_cgos_c_time_langinfo))
		break
	case int32(4):
		if idx > int32(0) {
			return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		}
		str = (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		break
	case int32(5):
		if idx > int32(3) {
			return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		}
		str = (*int8)(unsafe.Pointer(&_cgos_c_messages_langinfo))
		break
	default:
		return (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
	}
	for ; idx != 0; func() *int8 {
		idx--
		return func() (_cgo_ret *int8) {
			_cgo_addr := &str
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}() {
		for ; *str != 0; *(*uintptr)(unsafe.Pointer(&str))++ {
		}
	}
	if cat != int32(1) && int32(*str) != 0 {
		str = __lctrans(str, *(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&loc.Cat)))) + uintptr(cat)*8)))
	}
	return (*int8)(unsafe.Pointer(str))
}
func __nl_langinfo(item int32) *int8 {
	return __nl_langinfo_l(item, __pthread_self().Locale)
}
