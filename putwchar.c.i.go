package libc

func putwchar(c int32) uint32 {
	return fputwc(c, &__stdout_FILE)
}
