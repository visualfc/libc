package libc

func Wctomb(s *int8, wc int32) int32 {
	if !(s != nil) {
		return int32(0)
	}
	return int32(wcrtomb(s, wc, nil))
}
