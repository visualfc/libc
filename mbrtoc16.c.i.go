package libc

import unsafe "unsafe"

func mbrtoc16(pc16 *uint16, s *int8, n uint64, ps *Struct___mbstate_t) uint64 {
	if !(ps != nil) {
		ps = (*Struct___mbstate_t)(unsafe.Pointer(&_cgos_mbrtoc16_internal_state_mbrtoc16))
	}
	var pending *uint32 = (*uint32)(unsafe.Pointer(ps))
	if !(s != nil) {
		return mbrtoc16(nil, (*int8)(unsafe.Pointer(&[1]int8{'\x00'})), uint64(1), ps)
	}
	if int32(*pending) > int32(0) {
		if pc16 != nil {
			*pc16 = uint16(*pending)
		}
		*pending = uint32(0)
		return uint64(18446744073709551613)
	}
	var wc int32
	var ret uint64 = mbrtowc(&wc, s, n, ps)
	if ret <= uint64(4) {
		if wc >= int32(65536) {
			*pending = uint32(wc&int32(1023) + int32(56320))
			wc = int32(55232) + wc>>int32(10)
		}
		if pc16 != nil {
			*pc16 = uint16(wc)
		}
	}
	return ret
}

var _cgos_mbrtoc16_internal_state_mbrtoc16 uint32
