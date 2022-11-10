package libc

func Wcstombs(s *int8, ws *int32, n uint64) uint64 {
	return wcsrtombs(s, &[]*int32{ws}[0], n, nil)
}
