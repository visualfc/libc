package libc

func Nl_langinfo(item int32) *int8 {
	return __nl_langinfo(item)
}
