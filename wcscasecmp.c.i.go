package libc

func Wcscasecmp(l *int32, r *int32) int32 {
	return Wcsncasecmp(l, r, uint64(18446744073709551615))
}
