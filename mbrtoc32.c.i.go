package libc

import unsafe "unsafe"

func mbrtoc32(pc32 *uint32, s *int8, n uint64, ps *Struct___mbstate_t) uint64 {
	if !(ps != nil) {
		ps = (*Struct___mbstate_t)(unsafe.Pointer(&_cgos_mbrtoc32_internal_state_mbrtoc32))
	}
	if !(s != nil) {
		return mbrtoc32(nil, (*int8)(unsafe.Pointer(&[1]int8{'\x00'})), uint64(1), ps)
	}
	var wc int32
	var ret uint64 = mbrtowc(&wc, s, n, ps)
	if ret <= uint64(4) && pc32 != nil {
		*pc32 = uint32(wc)
	}
	return ret
}

var _cgos_mbrtoc32_internal_state_mbrtoc32 uint32
