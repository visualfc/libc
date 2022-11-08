package libc

import unsafe "unsafe"

func c16rtomb(s *int8, c16 uint16, ps *Struct___mbstate_t) uint64 {
	if !(ps != nil) {
		ps = (*Struct___mbstate_t)(unsafe.Pointer(&_cgos_c16rtomb_internal_state_c16rtomb))
	}
	var x *uint32 = (*uint32)(unsafe.Pointer(ps))
	var wc int32
	if !(s != nil) {
		if *x != 0 {
			goto ilseq
		}
		return uint64(1)
	}
	if !(*x != 0) && uint32(c16)-uint32(55296) < uint32(1024) {
		*x = uint32((int32(c16) - int32(55232)) << int32(10))
		return uint64(0)
	}
	if *x != 0 {
		if uint32(c16)-uint32(56320) >= uint32(1024) {
			goto ilseq
		} else {
			wc = int32(*x + uint32(c16) - uint32(56320))
		}
		*x = uint32(0)
	} else {
		wc = int32(c16)
	}
	return wcrtomb(s, wc, nil)
ilseq:
	*x = uint32(0)
	*X__errno_location() = int32(84)
	return uint64(18446744073709551615)
}

var _cgos_c16rtomb_internal_state_c16rtomb uint32
