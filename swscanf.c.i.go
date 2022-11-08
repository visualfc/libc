package libc

func swscanf(s *int32, fmt *int32, __cgo_args ...interface {
}) int32 {
	var ret int32
	var ap []interface {
	}
	ap = __cgo_args
	ret = vswscanf(s, fmt, ap)
	return ret
}
