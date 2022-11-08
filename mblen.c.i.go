package libc

func Mblen(s *int8, n uint64) int32 {
	return Mbtowc(nil, s, n)
}
