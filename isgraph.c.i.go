package libc

func Isgraph(c int32) int32 {
	return func() int32 {
		if uint32(c)-uint32(33) < uint32(94) {
			return 1
		} else {
			return 0
		}
	}()
}
func __isgraph_l(c int32, l *Struct___locale_struct) int32 {
	return Isgraph(c)
}
