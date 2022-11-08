package libc

func vwscanf(fmt *int32, ap []interface {
}) int32 {
	return vfwscanf(&__stdin_FILE, fmt, ap)
}
