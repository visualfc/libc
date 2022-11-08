package libc

func vwprintf(fmt *int32, ap []interface {
}) int32 {
	return vfwprintf(&__stdout_FILE, fmt, ap)
}
