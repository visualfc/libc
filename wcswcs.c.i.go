package libc

func Wcswcs(haystack *int32, needle *int32) *int32 {
	return Wcsstr(haystack, needle)
}
