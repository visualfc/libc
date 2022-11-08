package libc

func Wcsncasecmp_l(l *int32, r *int32, n uint64, locale *Struct___locale_struct) int32 {
	return Wcsncasecmp(l, r, n)
}
