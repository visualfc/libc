package libc

func c32rtomb(s *int8, c32 uint32, ps *Struct___mbstate_t) uint64 {
	return wcrtomb(s, int32(c32), ps)
}
