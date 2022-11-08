package libc

func putwc(c int32, f *Struct__IO_FILE) uint32 {
	return fputwc(c, f)
}
