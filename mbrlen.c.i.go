package libc

import unsafe "unsafe"

func mbrlen(s *int8, n uint64, st *Struct___mbstate_t) uint64 {
	return mbrtowc(nil, s, n, func() *Struct___mbstate_t {
		if st != nil {
			return st
		} else {
			return (*Struct___mbstate_t)(unsafe.Pointer(&_cgos_mbrlen_internal_mbrlen))
		}
	}())
}

var _cgos_mbrlen_internal_mbrlen uint32
