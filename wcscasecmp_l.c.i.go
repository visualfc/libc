package libc

func Wcscasecmp_l(l *int32, r *int32, locale *Struct___locale_struct) int32 {
	return Wcscasecmp(l, r)
}
