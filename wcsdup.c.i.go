package libc

func Wcsdup(s *int32) *int32 {
	var l uint64 = Wcslen(s)
	var d *int32 = (*int32)(Malloc((l + uint64(1)) * 4))
	if !(d != nil) {
		return (*int32)(nil)
	}
	return wmemcpy(d, s, l+uint64(1))
}
