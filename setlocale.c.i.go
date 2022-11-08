package libc

import unsafe "unsafe"

var _cgos_buf_setlocale [144]int8

func Setlocale(cat int32, name *int8) *int8 {
	var lm *struct___locale_map
	if uint32(cat) > uint32(6) {
		return (*int8)(nil)
	}
	__lock((*int32)(unsafe.Pointer(&__locale_lock)))
	if cat == int32(6) {
		var i int32
		if name != nil {
			var tmp_locale Struct___locale_struct
			var part [24]int8 = [24]int8{'C', '.', 'U', 'T', 'F', '-', '8', '\x00'}
			var p *int8 = name
			for i = int32(0); i < int32(6); i++ {
				var z *int8 = __strchrnul(p, ';')
				if uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(p)) <= uintptr(int64(23)) {
					Memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&part))), unsafe.Pointer(p), uint64(uintptr(unsafe.Pointer(z))-uintptr(unsafe.Pointer(p))))
					*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&part)))) + (uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(p))))) = int8(0)
					if *z != 0 {
						p = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(z)) + uintptr(int32(1))))
					}
				}
				lm = __get_locale(i, (*int8)(unsafe.Pointer(&part)))
				if uintptr(unsafe.Pointer(lm)) == uintptr(unsafe.Pointer((*struct___locale_map)(unsafe.Pointer(uintptr(18446744073709551615))))) {
					__unlock((*int32)(unsafe.Pointer(&__locale_lock)))
					return (*int8)(nil)
				}
				*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&tmp_locale.Cat)))) + uintptr(i)*8)) = lm
			}
			__libc.global_locale = tmp_locale
		}
		var s *int8 = (*int8)(unsafe.Pointer(&_cgos_buf_setlocale))
		var part *int8
		var same int32 = int32(0)
		for i = int32(0); i < int32(6); i++ {
			var lm *struct___locale_map = *(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__libc.global_locale.Cat)))) + uintptr(i)*8))
			if uintptr(unsafe.Pointer(lm)) == uintptr(unsafe.Pointer(*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__libc.global_locale.Cat)))) + uintptr(int32(0))*8)))) {
				same++
			}
			part = func() *int8 {
				if lm != nil {
					return (*int8)(unsafe.Pointer(&lm.name))
				} else {
					return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&[2]int8{'C', '\x00'}))))
				}
			}()
			var l uint64 = Strlen(part)
			Memcpy(unsafe.Pointer(s), unsafe.Pointer(part), l)
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(l))) = int8(';')
			*(*uintptr)(unsafe.Pointer(&s)) += uintptr(l + uint64(1))
		}
		*func() (_cgo_ret *int8) {
			_cgo_addr := &s
			*(*uintptr)(unsafe.Pointer(_cgo_addr))--
			return *_cgo_addr
		}() = int8(0)
		__unlock((*int32)(unsafe.Pointer(&__locale_lock)))
		return func() *int8 {
			if same == int32(6) {
				return (*int8)(unsafe.Pointer(part))
			} else {
				return (*int8)(unsafe.Pointer(&_cgos_buf_setlocale))
			}
		}()
	}
	if name != nil {
		lm = __get_locale(cat, name)
		if uintptr(unsafe.Pointer(lm)) == uintptr(unsafe.Pointer((*struct___locale_map)(unsafe.Pointer(uintptr(18446744073709551615))))) {
			__unlock((*int32)(unsafe.Pointer(&__locale_lock)))
			return (*int8)(nil)
		}
		*(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__libc.global_locale.Cat)))) + uintptr(cat)*8)) = lm
	} else {
		lm = *(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&__libc.global_locale.Cat)))) + uintptr(cat)*8))
	}
	var ret *int8 = func() *int8 {
		if lm != nil {
			return (*int8)(unsafe.Pointer((*int8)(unsafe.Pointer(&lm.name))))
		} else {
			return (*int8)(unsafe.Pointer(&[2]int8{'C', '\x00'}))
		}
	}()
	__unlock((*int32)(unsafe.Pointer(&__locale_lock)))
	return ret
}
